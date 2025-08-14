package main

import (
	"log"

	"github.com/kliment2000/go_kafka_pg/cache"
	"github.com/kliment2000/go_kafka_pg/consumer"
	"github.com/kliment2000/go_kafka_pg/db"
	"github.com/kliment2000/go_kafka_pg/server"
)

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	db.Init(dsn)

	if err := cache.Cache.LoadFromDB(); err != nil {
		log.Fatalf("failed to load cache: %v", err)
	}

	go consumer.StartKafkaConsumer()

	server.Start()
}
