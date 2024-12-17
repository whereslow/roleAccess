package control

import (
	"ValidStudio/global"
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

// ValidRole 验证对应token的身份
func ValidRole(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		Token    string `json:"token"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized", "auth": "NULL"})
		return
	}
	var flag = false
	role := req.Role
	token := req.Token
	username := req.Username
	flag = validate.Valid(username, token, role, global.RDB)
	c.JSON(200, gin.H{"flag": "success", "detail": "valid role success", "auth": flag})
	return
}
