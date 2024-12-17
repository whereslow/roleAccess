package control

import (
	"ValidStudio/DAO"
	"ValidStudio/global"
	"ValidStudio/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// Delete 通过admin权限的token删除指定用户名的账户
func Delete(c *gin.Context) {

	// request json解析及其参数绑定
	var req struct {
		Token          string `json:"token"`
		DeleteUsername string `json:"delete_username"`
		OPUsername     string `json:"op_username"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{"flag": "fail", "detail": "request is not standardized"})
		return
	}
	token := req.Token
	username := req.DeleteUsername
	OPUser := req.OPUsername
	isAdmin := validate.Valid(OPUser, token, "admin", global.RDB)
	if !isAdmin {
		// token非管理员
		c.JSON(200, gin.H{"flag": "fail", "detail": "you are not admin"})
		return
	} else {
		flag := DAO.DeleteUser(username, global.DB)
		if flag {
			slog.Info(fmt.Sprintf(" deleted success deleted user : %s, by admin user: %s", username, OPUser))
			c.JSON(200, gin.H{"flag": "success", "detail": "user deleted"})
			return
		} else {
			c.JSON(200, gin.H{"flag": "fail", "detail": "user does not exist"})
			return
		}
	}
}
