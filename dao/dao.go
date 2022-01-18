package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driver   = "mysql"
	username = "easypay"
	password = "123456"
	host     = "00000000"
	port     = "3306"
	database = "easypay"
	charset  = "utf8"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	DB, _ = sql.Open(driver, dsn)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		fmt.Println("err is", err)
		return
	}

	fmt.Println("连接数据库成功")
}
