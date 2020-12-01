package dao

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB gorm.DB

var (
	// DB .
	DB  *gorm.DB
	url = "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
)

func init() {
	DB = newDB()
}

func newDB() *gorm.DB {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatalf("database init: %v", err.Error())
	}
	db.DB().SetConnMaxLifetime(time.Duration(3) * time.Second)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(50)
	return db
}
