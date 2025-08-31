package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kliment2000/go_kafka_pg/cache"
	"github.com/kliment2000/go_kafka_pg/database"

	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "topic-A",
		GroupID:        "order-consumer",
		StartOffset:    kafka.LastOffset,
		CommitInterval: 0,
	})

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			log.Println("kafka read error:", err)
			continue
		}

		var parsed interface{}
		if err := json.Unmarshal(m.Value, &parsed); err != nil {
			log.Printf("json parse error for message %s: %v\n", m.Key, err)
			continue
		}

		orderMap, ok := parsed.(map[string]interface{})
		if !ok {
			log.Printf("invalid message format for order %s\n", m.Key)
			continue
		}

		orderUID, ok := orderMap["order_uid"].(string)
		if !ok {
			log.Printf("missing order_uid in message %s\n", m.Key)
			continue
		}

		tx, err := database.BeginTransaction()
		if err != nil {
			log.Printf("failed to begin transaction for order %s: %v\n", orderUID, err)
			continue
		}

		err = database.InsertOrderWithTx(tx, &database.Order{
			OrderUID: orderUID,
			Data:     m.Value,
		})
		if err != nil {
			log.Printf("database insert error for order %s: %v\n", orderUID, err)
			database.RollbackTransaction(tx)
			continue
		}

		cache.Cache.Put(orderUID, parsed)

		err = r.CommitMessages(context.Background(), m)
		if err != nil {
			log.Printf("failed to commit message %s: %v\n", m.Key, err)
			database.RollbackTransaction(tx)
			continue
		}

		database.CommitTransaction(tx)
		fmt.Printf("Order %s successfully processed\n", orderUID)
	}
}
