package config

import (
	"ValidStudio/global"
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

// InitRedis 创建成功无返回,创建失败返回Redis ping的err
func InitRedis() error {
	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	poolSize, _ := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
		PoolSize: poolSize,
	})

	ping := global.RDB.Ping(context.Background())
	if ping.Err() != nil {
		return ping.Err()
	} else {
		return nil
	}
}
