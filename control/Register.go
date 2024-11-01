package control

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"ValidStudio/validate"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	token := c.PostForm("token")
	newUsername := c.PostForm("new_username")
	newPassword := c.PostForm("new_password")
	newRole := c.PostForm("new_role")
	isAdmin := validate.ValidAdmin(token)
	if !isAdmin {
		c.JSON(200, gin.H{"fail": "you are not admin"})
		return
	} else {
		err := DAO.CreateUser(newUsername, newPassword, newRole, config.DB)
		if err != nil {
			c.JSON(200, gin.H{"fail": "server error"})
			return
		}
		c.JSON(200, gin.H{"success": "User created"})
		return
	}
}
