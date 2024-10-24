package test

import (
	"CasbinStudio/DAO"
	"CasbinStudio/config"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"math/rand/v2"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	// 我错了,不该不分离Service,以为东西少
	config.InitMysql()
	config.InitRedis()
	username := "lry"
	password := "www"
	role, isExist, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		return
	}
	if !isExist {
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
	time := time.Duration(2*3600000000000 + rand.IntN(10)*3600000000000) // 2-10小时
	_, err = config.RDB.Set(username, hashString, time).Result()         // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := config.RDB.Set(hashString, role, time).Result()
	if err != nil {
		slog.Error(err.Error())
	}
	if result == "OK" {
		slog.Info(fmt.Sprintf("login success, username: %s, role: %s", username, role))
	}
	// 返回token
}
func TestLoginOut(t *testing.T) {
	config.InitMysql()
	config.InitRedis()
	username := "lry"
	password := "www"
	_, isExist, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		return
	}
	if !isExist {
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
