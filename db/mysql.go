package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接初始化
func NewMySQLDB() (*sql.DB, error) {
	dsn := "root:12345@tcp(127.0.0.1:3306)/task_api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Ping() 真正确认连接可用
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("mysql connected")
	return db, nil
}
