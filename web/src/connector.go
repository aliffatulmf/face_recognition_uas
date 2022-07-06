package src

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open("root:@tcp/dumb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		os.Exit(1)
	}

	return
}
