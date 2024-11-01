package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_, isExist, err := DAO.AccessRole(username, password, config.DB)
	if err != nil {
		c.JSON(200, gin.H{"fail": "server error"})
		return
	}
	if !isExist {
		c.JSON(200, gin.H{"fail": "username or password error"})
		return
	}
	token := config.RDB.Get(username).Val()
	if token == "" {
		c.JSON(200, gin.H{"fail": "not login"})
	} else {
		// redis删除登录凭证
		token := config.RDB.Get(username).Val()
		config.RDB.Del(username)
		config.RDB.Del(token)
		c.JSON(200, gin.H{"success": "log out"})
	}
}
