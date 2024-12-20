package validate

import (
	"ValidStudio/global"
	"strings"
)

// Valid 验证token对应的身份
func Valid(username string, token string, toValidRole string) bool {
	roleList := strings.Split(toValidRole, "|")
	var role string
	var userToken string

	localRole, foundLocalRole := global.Cache.Get(token)
	if foundLocalRole {
		role = localRole.(string)
	}
	localUserToken, foundLocalUserToken := global.Cache.Get(username)
	if foundLocalUserToken {
		userToken = localUserToken.(string)
	}
	if role == "" {
		return false
	}
	if userToken != token {
		return false
	}
	for _, r := range roleList {
		// | 分隔的任意符合, -${role} 表示不接受该角色
		if r[0] == '-' && role == r[1:] {
			return false
		}
		if role == r {
			return true
		}
	}
	return false
}
