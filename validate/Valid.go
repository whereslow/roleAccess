package validate

import "ValidStudio/config"

// Valid 验证token对应的身份
func Valid(username string, token string, toValidRole string) bool {
	role := config.RDB.Get(token).Val()
	userToken := config.RDB.Get(username).Val()
	if role == "" {
		return false
	}
	return role == toValidRole && userToken == token
}
