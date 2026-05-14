package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(databaseDsn string) {

	// 4. Kết nối GORM
	db, err := gorm.Open(postgres.Open(databaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect Database failed: ", err)
	}

	fmt.Printf("Connect Database success")
	DB = db
}
