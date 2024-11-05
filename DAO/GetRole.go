package DAO

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

// AccessRole 根据账号和密码取得对应角色
// 返回值: (角色,函数调用是否成功,报错)
func AccessRole(username string, toValidPassword string, db *sqlx.DB) (string, bool, error) {
	sql := "SELECT role, password FROM roletable WHERE username = ?"
	selectSql := "SELECT username FROM roletable WHERE username=?"
	var role string
	var realPassword string
	var selected string
	err := db.Get(&selected, selectSql, username)
	if err != nil {
		slog.Error(err.Error())
	}
	if selected == "" {
		return "", false, nil
	}
	row := db.QueryRow(sql, username)
	err = row.Scan(&role, &realPassword)
	// sql错误
	if err != nil {
		slog.Error(err.Error())
		return "", false, err
	}
	// 用户不存在
	if role == "" || realPassword == "" {
		slog.Info(username + " " + toValidPassword + " ")
		return "", false, nil
	}
	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(toValidPassword))
	if err != nil {
		// 账号和密码不对应
		slog.Info(username + " " + toValidPassword + " not match")
		return "", false, nil
	} else {
		// 验证通过
		return role, true, nil
	}
}
