package repository

import (
	"context"
	"task-service/db"
	"task-service/model"
)

type TaskRepository struct{}

func (r *TaskRepository) Create(task *model.Task) error {
	query := `INSERT INTO tasks (title, status, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.DB.QueryRow(context.Background(), query,
		task.Title, task.Status, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
}
