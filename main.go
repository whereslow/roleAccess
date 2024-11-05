package main

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"ValidStudio/control"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		err = godotenv.Load("./.env")
		if err != nil {
			log.Fatal("not found .env file")
		}
	}
	err = config.InitMysql()
	if err != nil {
		panic(err)
	}
	err = config.InitRedis()
	if err != nil {
		panic(err)
	}

	// 插入初始admin
	DAO.CreateUser("whereslow", "whereslow", "admin", config.DB)
	// ~
	r := gin.Default()
	sso := r.Group("/sso")
	{
		sso.POST("/register", control.Register)
		sso.POST("/login", control.Login)
		sso.POST("/valid", control.ValidRole)
		sso.POST("/logout", control.LogOut)
		sso.DELETE("/delete", control.Delete)
	}

	err = r.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
