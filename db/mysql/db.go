package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//链接mysql数据库

var db *sqlx.DB

func InitDb() (err error) {
	// Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/auth_hood?parseTime=true"
	if db, err = sqlx.Open("mysql", dsn); err != nil {
		return
	}
	//设置连接池的最大链接数
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}
