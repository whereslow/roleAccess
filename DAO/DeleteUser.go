package DAO

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

// DeleteUser 通过指定用户名删除用户,删除成功返回true,删除失败返回false
func DeleteUser(username string, db *sqlx.DB) bool {
	// 检查用户是否存在
	selectSql := "SELECT username FROM roletable WHERE username=?"
	var selected string
	err := db.Get(&selected, selectSql, username)
	if err != nil {
		slog.Error(err.Error())
		return false
	}
	if selected == "" {
		// 不存在该用户
		return false
	} else {
		// 删除用户
		deleteSql := "DELETE FROM roletable WHERE username = ?"
		_, err = db.Exec(deleteSql, username)
		if err != nil {
			slog.Error(err.Error())
			return false
		}
		return true
	}

}
