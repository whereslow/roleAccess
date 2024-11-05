package test

import (
	"ValidStudio/config"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestInitRedis(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = config.InitRedis()
	if err != nil {
		log.Fatal(err)
	}
}
