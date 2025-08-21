package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"together_service/internal/di"
	"together_service/internal/handler"
)

func main() {
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
