package config

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// InitDB เชื่อมต่อฐานข้อมูล
func InitDB() {
	// โหลดค่าจากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dsn)
	log.Println("DATABASE_URL:", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.DB().SetMaxOpenConns(5)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Minute * 10)

	DB = db
	log.Println("Database connected!")
}
