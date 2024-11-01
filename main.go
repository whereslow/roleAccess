package main

import (
	"ValidStudio/config"
	"ValidStudio/control"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitMysql()
	if err != nil {
		panic(err)
	}
	err = config.InitRedis()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	sso := r.Group("/sso")
	{
		sso.POST("/register", control.Register)
		sso.POST("/login", control.Login)
		sso.POST("/valid", control.ValidRole)
		sso.POST("/logout", control.LogOut)
	}

	err = r.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
