package model

type KafkaTaskMessage struct {
    Type string    `json:"type"`
    Task Task      `json:"task"`
}
