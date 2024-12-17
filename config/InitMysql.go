package config

import (
	"ValidStudio/global"
	"github.com/jmoiron/sqlx"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

// InitMysql 创建成功无返回,创建失败返回sqlx的err
func InitMysql() error {
	global.DB, _ = sqlx.Open(os.Getenv("SQL_DRIVER"), os.Getenv("SQL_DATA_SOURCE"))
	if err := global.DB.Ping(); err != nil {
		return err
	}
	return nil
}
