package handler

import (
	"net/http"
	"strings"
)

type Routes struct {
	taskHandler *TaskHandler
}

func BuildRoutes(taskHandler *TaskHandler) *Routes {
	return &Routes{
		taskHandler: taskHandler,
	}
}

func (r *Routes) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Обработчик для всех маршрутов /tasks
	mux.HandleFunc("/tasks", r.handleTasks)

	return mux
}

func (r *Routes) handleTasks(w http.ResponseWriter, req *http.Request) {
	// Убираем /tasks из пути для получения оставшейся части
	path := strings.TrimPrefix(req.URL.Path, "/tasks")

	switch {
	case req.Method == "POST" && path == "":
		// POST /tasks - создание задачи
		r.taskHandler.CreateTaskHandler(w, req)
	case req.Method == "GET" && path == "":
		// GET /tasks - получение всех задач
		r.taskHandler.GetAllTasksHandler(w, req)
	case req.Method == "GET" && strings.HasPrefix(path, "/"):
		// GET /tasks/{id} - получение задачи по ID
		r.taskHandler.GetTaskByIDHandler(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
