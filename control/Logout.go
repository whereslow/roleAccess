package control

import (
	"ValidStudio/DAO"
	"ValidStudio/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

// LogOut 验证账号密码,退出登录,销毁token
func LogOut(c *gin.Context) {
	// request json解析及其参数绑定
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized"})
		return
	}
	username := req.Username
	password := req.Password

	ip := c.ClientIP()

	// ip 时间尝试限制
	ipCount, err := global.RDB.Get(global.Background, fmt.Sprintf("count_%s", ip)).Int()
	if err != nil && err.Error() != "nil" {
		slog.Error(err.Error())
	}
	if ipCount > 1000 {
		c.JSON(200, gin.H{"flag": "fail", "detail": "try too fast, please wait", "token": "NULL"})
		return
	}

	role, finish, err := DAO.AccessRole(username, password, global.DB)
	if err != nil {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "internal server error"})
		return
	}
	if !finish {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "username or password error"})
		return
	}
	token := global.RDB.Get(global.Background, username).Val()
	if token == "" {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "not login"})
		return
	} else {
		// 发送消息
		ctx, cancel := context.WithTimeout(global.Background, 30*time.Second)
		defer cancel()
		global.RDB.Publish(ctx, "logout", username)
		// 删除本地缓存
		global.Cache.Delete(username)
		global.Cache.Delete(token)
		slog.Info(fmt.Sprintf("logout success, username: %s, role: %s", username, role))
		c.JSON(200, gin.H{"flag": "success", "detail": "log out"})
		return
	}
}
