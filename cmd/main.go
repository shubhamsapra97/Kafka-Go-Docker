package main

import (
	"net/http"
	"task-service/db"
	"task-service/handler"
	"task-service/repository"
	"task-service/service"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	db.SetupDb()

	repo := &repository.TaskRepository{}
	svc := &service.TaskService{Repo: repo}
	h := &handler.TaskHandler{Service: svc}

	r := mux.NewRouter()

	r.HandleFunc("/createTask", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
