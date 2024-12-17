package main

import (
	"ValidStudio/DAO"
	"ValidStudio/config"
	"ValidStudio/control"
	"ValidStudio/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		gin.SetMode(gin.DebugMode)
		log.Println("you are on debug mode")
		err = godotenv.Load("./.env")
		if err != nil {
			log.Fatal("could not found .env file")
		}
	} else {
		if os.Args[1] != "deploy" {
			log.Println("with arg deploy to enable release mode")
			return
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	}
	err = config.InitMysql()
	if err != nil {
		log.Fatal("Mysql连接失败")
	}
	err = config.InitRedis()
	if err != nil {
		log.Fatal("Redis连接失败")
	}

	// 初始化缓存
	config.InitCache()

	// 插入初始admin, 如果存在用户则不会创建
	DAO.CreateUser("whereslow", "whereslow", "admin", global.DB)
	// ~
	r := gin.Default()
	// 跨域中间件
	r.Use(cors.Default())
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
		log.Println("服务器启动失败")
		panic(err)
	}
}
