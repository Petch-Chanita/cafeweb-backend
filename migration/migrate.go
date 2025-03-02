package migration

import (
	"cafeweb-backend/models"
	"log"

	"github.com/jinzhu/gorm"
)

func RunMigration(db *gorm.DB) {
	modelsToMigrate := []interface{}{
		&models.Cafes{},
		&models.Users{},
		&models.Products{},
		&models.Abouts{},
		&models.Images{},
		&models.Categories{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model).Error; err != nil {
			log.Fatalf("❌ Migration failed for %T: %v", model, err)
		}
	}

	log.Println("✅ All migrations completed successfully!")
}
