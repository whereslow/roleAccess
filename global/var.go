package global

import (
	redis2 "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

var RDB *redis2.Client
