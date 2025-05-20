package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func InitDB() {
	var err error
	connStr := "postgres://user:password@localhost:5432/taskdb"
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected to database")
}

func SetupDb() {
	fmt.Println("Setting up database...")
    query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );`
    _, err := DB.Exec(context.Background(), query)
    if err != nil {
        log.Fatalf("Failed to run migration: %v\n", err)
    }
	fmt.Println("Database setup complete")
}
