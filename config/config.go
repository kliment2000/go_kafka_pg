package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string
	ServerPort int
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден")
	}

	port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return &Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: port,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBName:     os.Getenv("DB_NAME"),
	}
}
