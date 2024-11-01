package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"testing"
)

func TestGetRole(t *testing.T) {
	config.InitMysql()
	DAO.AccessRole("lry", "www", config.DB)
}
