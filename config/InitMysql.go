package config

import (
	"github.com/jmoiron/sqlx"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

// DB mysql的全局变量
var DB *sqlx.DB

// InitMysql 创建成功无返回,创建失败返回sqlx的err
func InitMysql() error {
	DB, _ = sqlx.Open(os.Getenv("SQL_DRIVER"), os.Getenv("SQL_DATA_SOURCE"))
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}
