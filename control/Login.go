package control

import (
	"ValidStudio/DAO"
	"ValidStudio/global"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"math/rand/v2"
	"time"
)

// Login 登录,校验账号密码,返回临时token,并存储token和账户在redis,用于校验.如已经登录,则返回失败
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password is empty", "token": "NULL"})
		return
	}
	// RDB 设置ip : 账户
	ip := c.ClientIP()
	if global.RDB.Get(ip).Val() == username {
		c.JSON(200, gin.H{"flag": "fail", "detail": "login to fast", "token": "NULL"})
		return
	}
	global.RDB.Set(ip, username, 3*time.Second)
	// mysql 取值
	role, finish, err := DAO.AccessRole(username, password, global.DB)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "internal server error", "token": "NULL"})
		return
	}
	if !finish {
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password error", "token": "NULL"})
		return
	}
	// redis取值,验证是否登录,避免多token申请
	token := global.RDB.Get(username)
	if token.Val() != "" {
		c.JSON(200, gin.H{"flag": "fail", "detail": "have get a token , Can not get more token", "token": "NULL"})
		return
	}
	id, _ := uuid.NewRandom()
	hashId := sha256.Sum256([]byte(id.String()))
	hashStringToken := hex.EncodeToString(hashId[:])
	// redis存值
	t := time.Duration(2*3600000000000 + rand.IntN(1000)*36000000000) // 2-10小时
	_, err = global.RDB.Set(username, hashStringToken, t).Result()    // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := global.RDB.Set(hashStringToken, role, t).Result()
	if err != nil {
		slog.Error(err.Error())
	}
	if result == "OK" {
		slog.Info(fmt.Sprintf("login success, username: %s, role: %s", username, role))
	}
	// 返回token
	c.JSON(200, gin.H{"flag": "success", "detail": "have get token", "token": hashStringToken})
	return
}
