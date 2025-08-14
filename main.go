package main

import (
	"log"
	"wb_l01/cache"
	"wb_l01/consumer"
	"wb_l01/db"
	"wb_l01/server"
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
