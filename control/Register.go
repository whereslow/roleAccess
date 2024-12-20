package control

import (
	"ValidStudio/DAO"
	"ValidStudio/global"
	"ValidStudio/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

// Register 通过账号和密码,验证admin权限,有权限则创建对应身份的账户,储存在mysql
func Register(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Token       string `json:"token"`
		Username    string `json:"username"`
		NewUsername string `json:"new_username"`
		NewPassword string `json:"new_password"`
		NewRole     string `json:"new_role"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		slog.Error("", err.Error())
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized"})
		return
	}
	token := req.Token
	newUsername := req.NewUsername
	newPassword := req.NewPassword
	newRole := req.NewRole
	username := req.Username

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

	isAdmin := validate.Valid(username, token, "admin")
	if !isAdmin {
		global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
		global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
		c.JSON(200, gin.H{"flag": "fail", "detail": "you are not admin"})
		return
	} else {
		flag := DAO.CreateUser(newUsername, newPassword, newRole, global.DB)
		if !flag {
			global.RDB.Incr(global.Background, fmt.Sprintf("count_%s", ip))
			global.RDB.Expire(global.Background, fmt.Sprintf("count_%s", ip), 10*time.Minute)
			c.JSON(200, gin.H{"flag": "fail", "detail": "User has been Created"})
			return
		}
		c.JSON(200, gin.H{"flag": "success", "detail": "User created"})
		return
	}
}
