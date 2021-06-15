package config

import (
	"fmt"
	"log"

	"github.com/TomatoPals/Pomoduck-Go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Load app.env congfiguration then establish DB connection with env credentials
func InitialMigration() {

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DB, err = gorm.Open(mysql.Open(config.DB_CONNECTION_DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate((&models.User{}))
}
