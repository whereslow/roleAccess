package DAO

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

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
		// 验证不通过
		slog.Info(username + " " + toValidPassword + " not match")
		return "", true, err
	} else {
		// 验证通过
		return role, true, nil
	}
}
