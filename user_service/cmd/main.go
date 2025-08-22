package main

import (
	"log"
	"net/http"

	"user_service/internal/di"
	"user_service/internal/handler"

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
	userHandler := handler.NewUserHandler(container.ServiceContainer.UserService)

	// Настраиваем маршруты
	routes := handler.SetupRoutes(userHandler)

	// Получаем конфигурацию
	cfg := di.GetConfig()

	log.Printf("User service running on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, routes); err != nil {
		log.Fatal(err)
	}
}
