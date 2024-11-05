package control

import (
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

// ValidRole 验证对应token的身份
func ValidRole(c *gin.Context) {
	var flag = false
	role := c.PostForm("role")
	token := c.PostForm("token")
	flag = validate.Valid(token, role)

	c.JSON(200, gin.H{"success": "valid role success", "auth": flag})
	return
}
