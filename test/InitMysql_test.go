package test

import (
	"ValidStudio/config"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestInitMysql(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = config.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
}
