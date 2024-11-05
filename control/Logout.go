package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"github.com/gin-gonic/gin"
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
		c.JSON(200, gin.H{"fail": "request is not standardized"})
	}
	username := req.Username
	password := req.Password
	_, finish, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		c.JSON(200, gin.H{"fail": "internal server error"})
		return
	}
	if !finish {
		c.JSON(200, gin.H{"fail": "username or password error"})
		return
	}
	token := config.RDB.Get(username).Val()
	if token == "" {
		c.JSON(200, gin.H{"fail": "not login"})
		return
	} else {
		// redis删除登录凭证
		token := config.RDB.Get(username).Val()
		config.RDB.Del(username)
		config.RDB.Del(token)
		c.JSON(200, gin.H{"success": "log out"})
		return
	}
}
