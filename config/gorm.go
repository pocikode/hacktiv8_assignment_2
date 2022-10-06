package config

import (
	"assignment2/gin/repository"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBORM *gorm.DB

const (
	host   = "127.0.0.1"
	port   = "3306"
	user   = "root"
	pass   = "root"
	dbname = "hacktiv8_assignment_2"
)

func ConnectGorm() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Migrate(db)
	DBORM = db

	return nil
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&repository.Order{}, &repository.Item{})
	if err != nil {
		panic(err)
	}
}

func GetGorm() *gorm.DB {
	return DBORM
}
