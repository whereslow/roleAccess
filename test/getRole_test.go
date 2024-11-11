package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetRole(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitMysql()

	_, _, err = DAO.AccessRole("lry", "www", config.DB)
	if err != nil {
		log.Fatal(err)
	}
}
