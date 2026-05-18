package database

import (
	"be-summer-store/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// 4. Kết nối GORM
	db, err := gorm.Open(postgres.Open(config.AppConfig.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect Database failed: ", err)
	}

	fmt.Printf("Connect Database success")
	DB = db
}
