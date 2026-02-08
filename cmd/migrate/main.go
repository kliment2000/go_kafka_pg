package main

import (
	"fmt"

	"github.com/kliment2000/go_kafka_pg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kliment2000/go_kafka_pg/database"
)

func main() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(db); err != nil {
		panic(err)
	}
}
