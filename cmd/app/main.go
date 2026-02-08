package main

import (
	"fmt"
	"log"

	"github.com/kliment2000/go_kafka_pg/cache"
	"github.com/kliment2000/go_kafka_pg/config"
	"github.com/kliment2000/go_kafka_pg/consumer"
	"github.com/kliment2000/go_kafka_pg/database"
	"github.com/kliment2000/go_kafka_pg/server"
)

func main() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	database.Init(dsn)

	if err := cache.Cache.LoadFromDB(); err != nil {
		log.Fatalf("failed to load cache: %v", err)
	}

	go consumer.StartKafkaConsumer()

	server.Start()
}
