package services

import (
	"errors"

	"cafeweb-backend/models"

	"github.com/jinzhu/gorm"
)

// AboutService - struct สำหรับจัดการเกี่ยวกับ Abouts
type AboutService struct {
	DB *gorm.DB
}

// NewAboutService - ฟังก์ชันสร้าง AboutService
func NewAboutService(db *gorm.DB) *AboutService {
	return &AboutService{DB: db}
}

// CreateAbout - สร้างข้อมูล About ใหม่
func (s *AboutService) CreateAbout(about *models.Abouts) error {
	tx := s.DB.Begin()

	var existing models.Abouts

	// ค้นหา About ที่มี cafe_id เดียวกัน
	if err := tx.Where("cafe_id = ?", about.CafeID).First(&existing).Error; err == nil {
		// ถ้ามีอยู่แล้ว ให้ Update แทน
		if err := tx.Model(&existing).Updates(about).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// ถ้าไม่พบ ให้สร้างใหม่
		if err := tx.Create(about).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Error อื่น ๆ
		tx.Rollback()
		return err
	}

	// โหลดข้อมูล Cafe ที่เชื่อมโยง
	if err := tx.Preload("Cafe").First(about, "id = ?", about.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// UpdateAbout - อัปเดตข้อมูล About
func (s *AboutService) UpdateAbout(id string, updatedData *models.Abouts) error {
	var about models.Abouts
	if err := s.DB.First(&about, "id = ?", id).Error; err != nil {
		return errors.New("record not found")
	}

	if updatedData.AboutEn != nil {
		about.AboutEn = updatedData.AboutEn
	}
	if updatedData.AboutTh != nil {
		about.AboutTh = updatedData.AboutTh
	}
	if updatedData.ImageURL != "" {
		about.ImageURL = updatedData.ImageURL
	}

	if err := s.DB.Save(&about).Error; err != nil {
		return err
	}
	return nil
}

func (s *AboutService) GetAboutByCafeID(cafeID string) (*models.Abouts, error) {
	var about models.Abouts

	// ดึงข้อมูล About โดยใช้ cafe_id
	err := s.DB.Preload("Cafe").Where("cafe_id = ?", cafeID).First(&about).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("About not found for the given cafe_id")
		}
		return nil, err // หากเกิด error อื่นๆ
	}

	return &about, nil
}
