package test

import (
	"ValidStudio/config"
	"ValidStudio/validate"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestValidateAdmin(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = config.InitRedis()
	if err != nil {
		panic(err)
	}
	token := config.RDB.Get("lry").Val()
	flag := validate.Valid(token, "admin")
	if !flag {
		t.Error("Admin validation failed")
	}
}
