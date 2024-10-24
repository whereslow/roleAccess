package test

import (
	"CasbinStudio/DAO"
	"CasbinStudio/config"
	"testing"
)

func TestCreateUser(t *testing.T) {
	config.InitMysql()
	DAO.CreateUser("lry", "www", "admin", config.DB)
}
