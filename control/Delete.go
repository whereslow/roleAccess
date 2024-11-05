package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

// Delete 通过admin权限的token删除指定用户名的账户
func Delete(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Token          string `json:"token"`
		DeleteUsername string `json:"delete_username"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"fail": "request is not standardized"})
	}
	token := req.Token
	username := req.DeleteUsername
	isAdmin := validate.Valid(token, "admin")
	if !isAdmin {
		// token非管理员
		c.JSON(200, gin.H{"fail": "you are not admin"})
		return
	} else {
		flag := DAO.DeleteUser(username, config.DB)
		if flag {
			c.JSON(200, gin.H{"success": "user deleted"})
		} else {
			c.JSON(200, gin.H{"fail": "user does not exist"})
		}
	}
}
