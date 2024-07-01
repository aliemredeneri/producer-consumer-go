package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "example-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()

	messages := make(chan string)

	go func() {
		for {
			messages <- "Hello Kafka"
			time.Sleep(10 * time.Millisecond)
		}
	}()

	for msg := range messages {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(msg)},
		)
		if err != nil {
			log.Println("failed to write message:", err)
		} else {
			log.Println("sent message:", msg)
		}
	}
}
