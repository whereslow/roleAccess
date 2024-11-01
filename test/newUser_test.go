package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"testing"
)

func TestCreateUser(t *testing.T) {
	config.InitMysql()
	DAO.CreateUser("lry", "www", "admin", config.DB)
}
