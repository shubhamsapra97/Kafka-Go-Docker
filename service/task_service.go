package service

import (
	"time"
	"task-service/model"
	"task-service/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) Create(task *model.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt
	return s.Repo.Create(task)
}

func (s *TaskService) GetAll(status string, limit, offset int) ([]model.Task, error) {
	return s.Repo.GetAll(status, limit, offset)
}
