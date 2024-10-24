package validate

import "CasbinStudio/config"

func ValidStudent(token string) bool {
	role := config.RDB.Get(token).Val()
	return role == "student"
}
