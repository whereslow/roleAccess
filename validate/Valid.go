package validate

import (
	"ValidStudio/global"
	redis2 "github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
	"strings"
)

// Valid 验证token对应的身份
func Valid(username string, token string, toValidRole string, rdb *redis2.Client) bool {
	roleList := strings.Split(toValidRole, "|")
	var role string
	var userToken string

	// 一级缓存和二级缓存
	tokenLocal, found := global.Cache.Get(token)
	if found {
		role = tokenLocal.(string)
	} else {
		role = rdb.Get(token).Val()
		global.Cache.Set(token, role, cache.DefaultExpiration)
	}
	userTokenLocal, found := global.Cache.Get(username)
	if found {
		userToken = userTokenLocal.(string)
	} else {
		userToken = rdb.Get(username).Val()
		global.Cache.Set(username, userToken, cache.DefaultExpiration)
	}

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
