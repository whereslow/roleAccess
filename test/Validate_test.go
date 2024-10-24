package test

import (
	"CasbinStudio/config"
	"CasbinStudio/validate"
	"testing"
)

func TestValidateAdmin(t *testing.T) {
	config.InitRedis()
	token := config.RDB.Get("lry").Val()
	flag := validate.ValidAdmin(token)
	if !flag {
		t.Error("Admin validation failed")
	}
}
