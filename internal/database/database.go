package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"studentPlanner/internal/models"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/plannerDB?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err := Db.AutoMigrate(&models.Student{}, &models.Plan{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
