package handler

import (
	"encoding/json"
	"net/http"
	"test_lo/internal/logger"
	"test_lo/internal/models"

	"github.com/go-chi/chi/v5"
)

type TaskService interface {
	CreateTask(title, description string) (models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	GetAllTasks(status string) ([]models.Task, error)
}

type TaskHandler struct {
	service TaskService
	logger  *logger.Logger
}

func MakeTaskHandler(s TaskService, l *logger.Logger) *TaskHandler {
	return &TaskHandler{service: s, logger: l}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newTask, err := h.service.CreateTask(task.Title, task.Description)
	if err != nil {
		h.logger.Log("ERROR", "Failed to create task: "+err.Error())
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	h.logger.Log("INFO", "Task created successfully: "+newTask.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		h.logger.Log("WARN", "Task not found: "+id)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.logger.Log("INFO", "Task retrieved successfully: "+id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	tasks, err := h.service.GetAllTasks(status)
	if err != nil {
		h.logger.Log("ERROR", "Failed to get all tasks: "+err.Error())
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	h.logger.Log("INFO", "Tasks retrieved successfully, count: "+string(rune(len(tasks))))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
