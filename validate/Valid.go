package validate

import "ValidStudio/config"

func Valid(token string, toValidRole string) bool {
	role := config.RDB.Get(token).Val()
	if role == "" {
		return false
	}
	return role == toValidRole
}
