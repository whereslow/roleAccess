package DAO

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func CreateUser(username string, password string, role string, db *sqlx.DB) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
	}
	// 查询是否已存在
	selectSql := "SELECT username FROM roletable WHERE username=?"
	var selected string
	err = db.Get(&selected, selectSql, username)
	if err != nil {
		slog.Error(err.Error())
	}
	if selected == "" {
		insertSql := "INSERT INTO roletable(roletable.username,roletable.`password`,roletable.role)VALUES(?,?,?)"
		db.MustExec(insertSql, username, string(hashPassword), role)
		return nil
	} else {
		slog.Info(username + " have exist but still to be register")
		return errors.New("exist user")
	}

}
