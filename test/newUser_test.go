package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitMysql()
	DAO.CreateUser("lry", "www", "admin", config.DB)
}
