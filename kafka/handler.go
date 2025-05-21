package kafka

import (
    "encoding/json"
    "fmt"
    "task-service/model"
    "task-service/service"
)

func handleTaskMessage(data []byte, svc *service.TaskService) {
    var msg model.KafkaTaskMessage
    if err := json.Unmarshal(data, &msg); err != nil {
        fmt.Printf("Error unmarshaling message: %v", err)
        return
    }

    switch msg.Type {
    case "create":
        if err := svc.Create(&msg.Task); err != nil {
            fmt.Printf("Error creating task: %v", err)
            return
        }
        fmt.Printf("Task created from Kafka: %+v\n", msg.Task)
    case "update":
        if err := svc.Update(&msg.Task); err != nil {
            fmt.Printf("Error updating task: %v", err)
            return
        }
        fmt.Printf("Task updated from Kafka: %+v\n", msg.Task)
    default:
        fmt.Printf("Unknown message type: %s\n", msg.Type)
    }
}
