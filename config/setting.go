package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
