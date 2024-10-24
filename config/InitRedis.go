package config

import (
	redis2 "github.com/go-redis/redis"
)

var RDB *redis2.Client

func InitRedis() error {
	RDB = redis2.NewClient(&redis2.Options{
		Addr:     "localhost:26379",
		Password: "123456",
		DB:       0,
		PoolSize: 1000,
	})

	ping := RDB.Ping()
	if ping.Err() != nil {
		return ping.Err()
	} else {
		return nil
	}
}
