package validate

import "CasbinStudio/config"

func ValidTeacher(token string) bool {
	role := config.RDB.Get(token).Val()
	return role == "teacher"
}
