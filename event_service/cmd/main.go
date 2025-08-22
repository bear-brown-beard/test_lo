package main

import (
	"event_service/internal/di"
	"event_service/internal/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из config.env
	if err := godotenv.Load("config.env"); err != nil {
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

	log.Printf("Event service running on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, routes); err != nil {
		log.Fatal(err)
	}
}
