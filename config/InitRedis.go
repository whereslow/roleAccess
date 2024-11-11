package config

import (
	redis2 "github.com/go-redis/redis"
	"os"
	"strconv"
)

// RDB redis的全局变量
var RDB *redis2.Client

// InitRedis 创建成功无返回,创建失败返回Redis ping的err
func InitRedis() error {
	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	poolSize, _ := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	RDB = redis2.NewClient(&redis2.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
		PoolSize: poolSize,
	})

	ping := RDB.Ping()
	if ping.Err() != nil {
		return ping.Err()
	} else {
		return nil
	}
}
