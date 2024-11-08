package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

// Register 通过账号和密码,验证admin权限,有权限则创建对应身份的账户,储存在mysql
func Register(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Token       string `json:"token"`
		NewUsername string `json:"new_username"`
		NewPassword string `json:"new_password"`
		NewRole     string `json:"new_role"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized"})
	}
	token := req.Token
	newUsername := req.NewUsername
	newPassword := req.NewPassword
	newRole := req.NewRole
	isAdmin := validate.Valid(token, "admin")
	if !isAdmin {
		c.JSON(200, gin.H{"flag": "fail", "detail": "you are not admin"})
		return
	} else {
		flag := DAO.CreateUser(newUsername, newPassword, newRole, config.DB)
		if !flag {
			c.JSON(200, gin.H{"flag": "fail", "detail": "User has been Created"})
			return
		}
		c.JSON(200, gin.H{"flag": "success", "detail": "User created"})
		return
	}
}
