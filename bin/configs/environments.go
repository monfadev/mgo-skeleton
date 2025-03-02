package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironments() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Environments success")
}
