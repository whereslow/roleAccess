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
	err = config.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = DAO.AccessRole("lry", "www", config.DB)
	if err != nil {
		log.Fatal(err)
	}
}
