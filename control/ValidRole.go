package control

import (
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

func ValidRole(c *gin.Context) {
	var flag = false
	role := c.PostForm("role")
	token := c.PostForm("token")
	switch role {
	case "admin":
		flag = validate.ValidAdmin(token)
	case "student":
		flag = validate.ValidStudent(token)
	case "teacher":
		flag = validate.ValidTeacher(token)
	default:
		c.JSON(200, gin.H{"fail": "invalid role value"})
		return
	}
	c.JSON(200, gin.H{"success": "valid role success", "auth": flag})
	return
}
