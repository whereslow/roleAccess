package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"math/rand/v2"
	"time"
)

// Login 登录,校验账号密码,返回临时token,并存储token和账户在redis,用于校验.如已经登录,则返回失败
func Login(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized", "token": "NULL"})
	}
	username := req.Username
	password := req.Password
	if username == "" || password == "" {
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password is empty", "token": "NULL"})
		return
	}
	// mysql 取值
	role, finish, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "internal server error", "token": "NULL"})
		return
	}
	if !finish {
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password error", "token": "NULL"})
		return
	}
	// redis取值,验证是否登录,避免多token申请
	token := config.RDB.Get(username)
	if token.Val() != "" {
		c.JSON(200, gin.H{"flag": "fail", "detail": "have get a token , Can not get more token", "token": "NULL"})
		return
	}
	id, _ := uuid.NewRandom()
	hashId := sha256.Sum256([]byte(id.String()))
	hashStringToken := base64.URLEncoding.EncodeToString(hashId[:])
	// redis存值
	t := time.Duration(2*3600000000000 + rand.IntN(10)*3600000000000) // 2-10小时
	_, err = config.RDB.Set(username, hashStringToken, t).Result()    // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := config.RDB.Set(hashStringToken, role, t).Result()
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
