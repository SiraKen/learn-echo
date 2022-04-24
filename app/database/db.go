package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:root@tcp(mysql:3306)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	return db
}
