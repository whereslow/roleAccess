package test

import (
	"CasbinStudio/DAO"
	"CasbinStudio/config"
	"testing"
)

func TestGetRole(t *testing.T) {
	config.InitMysql()
	DAO.AccessRole("lry", "www", config.DB)
}
