package config

import "github.com/jmoiron/sqlx"
import _ "github.com/go-sql-driver/mysql"

var DB *sqlx.DB

func InitMysql() error {
	DB, _ = sqlx.Open("mysql", "root:root@tcp(localhost:3306)/vilid")
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}
