package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"github.com/joho/godotenv"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	godotenv.Load("../.env")
	config.InitMysql()
	flag := DAO.DeleteUser("lrq", config.DB)
	if !flag {
		t.Log("不存在用户")
	} else {
		t.Log("删除成功")
	}
}
