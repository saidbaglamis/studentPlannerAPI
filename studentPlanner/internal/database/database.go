package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	models2 "studentPlanner/internal/models"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/SKETCH?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&models2.Student{}, &models2.Plan{})
}
