package kafka

import (
    "context"
    "fmt"
    "task-service/service"

    "github.com/segmentio/kafka-go"
)

func StartTaskConsumer(svc *service.TaskService) {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"kafka:9092"},
        Topic:   "tasks",
        GroupID: "task-service-consumer-group",
    })
    defer r.Close()

    fmt.Println("Kafka consumer started, listening for tasks...")

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            fmt.Printf("Error reading message: %v", err)
            continue
        }
        handleTaskMessage(m.Value, svc)
    }
}
