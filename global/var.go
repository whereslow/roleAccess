package global

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

var DB *sqlx.DB

var RDB *redis.Client

var Cache *cache.Cache

var Background = context.Background()
