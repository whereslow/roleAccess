package control

import (
	"CasbinStudio/DAO"
	"CasbinStudio/config"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"math/rand/v2"
	"time"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{"fail": "username or password is empty"})
		return
	}
	// mysql 取值
	role, isExist, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		c.JSON(200, gin.H{"fail": "server error"})
		return
	}
	if !isExist {
		c.JSON(200, gin.H{"fail": "username or password error"})
		return
	}
	// redis取值,验证是否登录,避免多token申请
	token := config.RDB.Get(username)
	if token.Val() != "" {
		c.JSON(200, gin.H{"fail": "have get a token , Can not get more token"})
		return
		// 后续应该改成返回redis的token
	}
	// ~
	id, _ := uuid.NewRandom()
	hashId := sha256.Sum256([]byte(id.String()))
	hashString := base64.URLEncoding.EncodeToString(hashId[:])
	// redis存值
	t := time.Duration(2*3600000000000 + rand.IntN(10)*3600000000000) // 2-10小时
	_, err = config.RDB.Set(username, hashString, t).Result()         // 登录表,防止用户换取多token
	if err != nil {
		slog.Error(err.Error())
	}
	result, err := config.RDB.Set(hashString, role, t).Result()
	if err != nil {
		slog.Error(err.Error())
	}
	if result == "OK" {
		slog.Info(fmt.Sprintf("login success, username: %s, role: %s", username, role))
	}
	// 返回token
	c.JSON(200, gin.H{"success": "have get token", "token": hashString})
}
