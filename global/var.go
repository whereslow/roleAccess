package global

import (
	redis2 "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	cache2 "github.com/patrickmn/go-cache"
)

var DB *sqlx.DB

var RDB *redis2.Client

var Cache *cache2.Cache
