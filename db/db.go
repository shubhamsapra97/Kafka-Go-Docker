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
