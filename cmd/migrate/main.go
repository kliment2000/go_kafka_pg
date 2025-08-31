package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kliment2000/go_kafka_pg/database"
)

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(db); err != nil {
		panic(err)
	}
}
