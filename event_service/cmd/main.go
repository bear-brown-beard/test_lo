package main

import (
	"context"
	"event_service/internal/di"
	"event_service/internal/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Загружаем переменные окружения из config.env
	if err := loadEnvFile("config.env"); err != nil {
		log.Printf("Предупреждение: не удалось загрузить config.env: %v", err)
	}

	// Создаем основной контейнер зависимостей
	container, err := di.NewMainContainer()
	if err != nil {
		log.Fatal("Ошибка создания контейнера зависимостей:", err)
	}
	defer container.Close()

	// Создаем обработчик
	dateHandler := handler.NewDateHandler(container.ServiceContainer.DateService)

	// Настраиваем маршруты
	routes := handler.SetupRoutes(dateHandler)

	// Получаем конфигурацию
	cfg := di.GetConfig()

	// Создаем HTTP сервер
	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: routes,
	}

	// Канал для graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем сервер в горутине
	go func() {
		log.Printf("Сервис запущен на порту %s", cfg.Server.Port)
		log.Printf("API доступен по адресу: http://localhost:%s/api", cfg.Server.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Ошибка запуска сервера:", err)
		}
	}()

	// Ждем сигнал для graceful shutdown
	<-done
	log.Println("Получен сигнал завершения, закрываем сервер...")

	// Даем серверу время на завершение текущих запросов
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Ошибка graceful shutdown: %v", err)
	}

	log.Println("Сервер успешно остановлен")
}

// loadEnvFile загружает переменные окружения из файла
func loadEnvFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}

	return nil
}
