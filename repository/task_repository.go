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

func (r *TaskRepository) GetAll(status string, limit, offset int) ([]model.Task, error) {
	query := `SELECT id, title, status, created_at, updated_at FROM tasks`
	args := []interface{}{}

	if status != "" {
        query += " WHERE status=$1 ORDER BY id LIMIT $2 OFFSET $3"
        args = append(args, status, limit, offset)
    } else {
        query += " ORDER BY id LIMIT $1 OFFSET $2"
        args = append(args, limit, offset)
    }

	rows, err := db.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) Update(task *model.Task) error {
	// Begin a transaction
    tx, err := db.DB.Begin(context.Background())
    if err != nil {
        return err
    }
    defer tx.Rollback(context.Background())

    // Lock the row for update
    row := tx.QueryRow(context.Background(), "SELECT id FROM tasks WHERE id=$1 FOR UPDATE", task.ID)
    var id int
    if err := row.Scan(&id); err != nil {
        return err // row not found or other error
    }

    // Perform the update
    _, err = tx.Exec(context.Background(),
        "UPDATE tasks SET title=$1, status=$2, updated_at=$3 WHERE id=$4",
        task.Title, task.Status, task.UpdatedAt, task.ID)
    if err != nil {
        return err
    }

	// Commit the transaction
    return tx.Commit(context.Background())
}
