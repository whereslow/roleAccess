package test

import (
	"ValidStudio/config"
	"ValidStudio/validate"
	"testing"
)

func TestValidateAdmin(t *testing.T) {
	err := config.InitRedis()
	if err != nil {
		panic(err)
	}
	token := config.RDB.Get("lry").Val()
	flag := validate.Valid(token, "admin")
	if !flag {
		t.Error("Admin validation failed")
	}
}
