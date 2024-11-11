package config

import (
	redis2 "github.com/go-redis/redis"
	"os"
	"strconv"
)

// RDB redis的全局变量
var RDB *redis2.Client

// InitRedis 创建成功无返回,创建失败返回Redis ping的err
func InitRedis() {
	dbNum, _ := strconv.Atoi(os.Getenv("redis_db"))
	poolSize, _ := strconv.Atoi(os.Getenv("redis_pool_size"))
	RDB = redis2.NewClient(&redis2.Options{
		Addr:     os.Getenv("redis_address"),
		Password: os.Getenv("redis_password"),
		DB:       dbNum,
		PoolSize: poolSize,
	})
}
