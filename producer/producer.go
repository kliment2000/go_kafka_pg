package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	data, err := os.ReadFile("producer/order.json")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "topic-A",
		AllowAutoTopicCreation: true,
	}

	messages := []kafka.Message{
		{
			Key:   []byte("order-key"),
			Value: data,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = w.WriteMessages(ctx, messages...)

	if err != nil {
		log.Fatalf("Ошибка при отправке: %v", err)
	} else {
		log.Println("JSON-файл отправлен в Kafka")
	}

	if err := w.Close(); err != nil {
		log.Fatal("Ошибка при закрытии writer:", err)
	}
}
