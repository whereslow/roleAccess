package control

import (
	"ValidStudio/DAO"
	"ValidStudio/global"
	"context"
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
	var err error
	if username == "" || password == "" {
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password is empty", "token": "NULL"})
		return
	}
	// RDB 设置ip : 账户
	ip := c.ClientIP()
	// ip:用户 时间尝试限制
	if global.RDB.Get(global.Background, ip).Val() == username {
		c.JSON(200, gin.H{"flag": "fail", "detail": "login too fast", "token": "NULL"})
		return
	}
	// ip 时间尝试限制
	ipCount, err := global.RDB.Get(global.Background, fmt.Sprintf("count_%s", ip)).Int()
	if err != nil && err.Error() != "nil" {
		slog.Error(err.Error())
	}
	if ipCount > 1000 {
		c.JSON(200, gin.H{"flag": "fail", "detail": "login too fast, , please wait", "token": "NULL"})
		return
	}
	global.RDB.Set(global.Background, ip, username, 3*time.Second)
	// mysql 取值
	role, finish, err := DAO.AccessRole(username, password, global.DB)
	if err != nil {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "internal server error", "token": "NULL"})
		return
	}
	if !finish {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password error", "token": "NULL"})
		return
	}
	// redis取值,验证是否登录,避免多token申请
	token := global.RDB.Get(global.Background, username)
	if token.Val() != "" {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "have get a token , Can not get more token", "token": "NULL"})
		return
	}
	id, _ := uuid.NewRandom()
	hashId := sha256.Sum256([]byte(id.String()))
	hashStringToken := hex.EncodeToString(hashId[:])
	// redis存值
	t := time.Duration(2*3600000000000 + rand.IntN(1000)*36000000000)
	_, err = global.RDB.Set(global.Background, username, hashStringToken, 30*time.Second).Result() // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := global.RDB.Set(global.Background, hashStringToken, role, 30*time.Second).Result()
	if err != nil {
		slog.Error(err.Error())
	}
	if result == "OK" {
		slog.Info(fmt.Sprintf("login success, username: %s, role: %s", username, role))
	}
	// 发布消息
	ctx, cancel := context.WithTimeout(global.Background, 30*time.Second)
	defer cancel()
	global.RDB.Publish(ctx, "login", username)
	// 设置本地缓存
	global.Cache.Set(username, hashStringToken, t)
	global.Cache.Set(hashStringToken, role, t)
	// 返回token
	c.JSON(200, gin.H{"flag": "success", "detail": "have get token", "token": hashStringToken})
	return
}
