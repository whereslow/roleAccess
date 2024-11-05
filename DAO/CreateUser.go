package DAO

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

// CreateUser 根据账号密码和角色创建账号
func CreateUser(username string, password string, role string, db *sqlx.DB) bool {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
	}
	// 查询是否已存在
	selectSql := "SELECT username FROM roletable WHERE username=?"
	var selected string
	_ = db.Get(&selected, selectSql, username)
	if selected == "" {
		insertSql := "INSERT INTO roletable(roletable.username,roletable.`password`,roletable.role)VALUES(?,?,?)"
		db.MustExec(insertSql, username, string(hashPassword), role)
		return true
	} else {
		slog.Info(username + " have exist but still to be register")
		return false
	}

}
