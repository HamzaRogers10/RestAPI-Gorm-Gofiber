package config

import (
	"GoFiber/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DNS = "root:Hamza@10@tcp(127.0.0.1:3306)/devdb?charset=utf8mb4&parseTime=True&loc=Local"

var err error

func Migration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	Developer := models.Developer{}
	err := DB.AutoMigrate(&Developer)
	if err != nil {
		return
	}
}
