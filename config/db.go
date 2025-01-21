package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// konfigurasi db
	dsn := "root:@tcp(localhost:3306)/db_go?parseTime=true"
	// koneksi ke db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failde to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database Connected")
	DB = db
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting SQL DB instance: %v", err)
	}

	sqlDB.Close()
	log.Printf("Database connection closed")
}
