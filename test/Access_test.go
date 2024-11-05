package test

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"math/rand/v2"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	// dotenv load
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = config.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	err = config.InitRedis()
	if err != nil {
		log.Fatal(err)
	}
	username := "lry"
	password := "www"
	role, finish, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		t.Fatal(err)
	}
	if !finish {
		return
	}
	// 验证是否登录,避免多token申请
	token := config.RDB.Get(username)
	if token.Val() != "" {
		println(token.Val())
		return
	}
	// ~
	id, _ := uuid.NewRandom()
	hashId := sha256.Sum256([]byte(id.String()))
	hashString := base64.URLEncoding.EncodeToString(hashId[:])
	// redis存值
	saveTime := time.Duration(2*3600000000000 + rand.IntN(10)*3600000000000) // 2-10小时
	_, err = config.RDB.Set(username, hashString, saveTime).Result()         // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := config.RDB.Set(hashString, role, saveTime).Result()
	if err != nil {
		slog.Error(err.Error())
	}
	if result == "OK" {
		slog.Info(fmt.Sprintf("login success, username: %s, role: %s", username, role))
	}
	// 返回token
}
func TestLoginOut(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = config.InitMysql()
	if err != nil {
		log.Fatal(err)
	}
	err = config.InitRedis()
	if err != nil {
		log.Fatal(err)
	}
	username := "lry"
	password := "www"
	_, finish, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		return
	}
	if !finish {
		return
	}
	token := config.RDB.Get(username).String()
	if token == "" {
		return
	} else {
		// redis删除登录凭证
		token := config.RDB.Get(username).Val()
		config.RDB.Del(username)
		config.RDB.Del(token)
	}
}
