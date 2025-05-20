package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-service/model"
	"task-service/service"

	"github.com/gorilla/mux"
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

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(offsetStr)

	tasks, err := h.Service.GetAll(status, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var t model.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.ID = id
	if err := h.Service.Update(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(t)
}
