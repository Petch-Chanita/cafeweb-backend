package migration

import (
	"cafeweb-backend/config"
	"cafeweb-backend/models"
	"log"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&models.Cafes{},      // ใช้ BaseModel ใน Cafe
		&models.Users{},      // ใช้ BaseModel ใน User
		&models.Products{},   // ใช้ BaseModel ใน Product
		&models.Abouts{},     // ใช้ BaseModel ใน About
		&models.Images{},     // ใช้ BaseModel ใน Image
		&models.Categories{}, // ใช้ BaseModel ใน Category

	)

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully!")
}
