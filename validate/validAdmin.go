package validate

import "CasbinStudio/config"

func ValidAdmin(token string) bool {
	role := config.RDB.Get(token).Val()
	return role == "admin"
}
