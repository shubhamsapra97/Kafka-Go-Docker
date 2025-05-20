package handler

import (
	"encoding/json"
	"net/http"
	"task-service/model"
	"task-service/service"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t model.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(t)
}
