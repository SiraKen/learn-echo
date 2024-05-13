package database

import (
	m "lecho/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/echo?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&m.User{})

	return db
}
