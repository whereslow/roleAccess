package validate

import (
	redis2 "github.com/go-redis/redis"
	"strings"
)

// Valid 验证token对应的身份
func Valid(username string, token string, toValidRole string, rdb *redis2.Client) bool {
	roleList := strings.Split(toValidRole, "|")
	var role string
	var userToken string

	role = rdb.Get(token).Val()
	userToken = rdb.Get(username).Val()

	if role == "" {
		return false
	}
	if userToken != token {
		return false
	}
	for _, r := range roleList {
		// 任意符合
		if role == r {
			return true
		}
	}
	return false
}
