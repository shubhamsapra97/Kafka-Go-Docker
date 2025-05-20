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

	http.ListenAndServe(":8080", r)
}
