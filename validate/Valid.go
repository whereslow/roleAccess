package validate

import "ValidStudio/config"

// Valid 验证token对应的身份
func Valid(token string, toValidRole string) bool {
	role := config.RDB.Get(token).Val()
	if role == "" {
		return false
	}
	return role == toValidRole
}
