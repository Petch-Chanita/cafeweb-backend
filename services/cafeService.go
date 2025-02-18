package services

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/models"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type CafeService struct {
	db *gorm.DB
}

// NewCafeService - ฟังก์ชันสำหรับสร้าง CafeService
func NewCafeService(db *gorm.DB) *CafeService {
	return &CafeService{
		db: db,
	}
}

func (s *CafeService) CreateCafe(req dto.CafeRequest) (out models.Cafes, err error) {
	tx := s.db.Begin()

	// แปลงค่า string เป็น time.Time สำหรับ OpeningTime และ ClosingTime
	openingTime, err := time.Parse("15:04:05", req.OpeningTime)
	if err != nil {
		tx.Rollback()
		return out, fmt.Errorf("invalid opening time format: %v", err)
	}

	closingTime, err := time.Parse("15:04:05", req.ClosingTime)
	if err != nil {
		tx.Rollback()
		return out, fmt.Errorf("invalid closing time format: %v", err)
	}

	// แปลง dto.CafeRequest เป็น models.Cafes
	cafe := models.Cafes{
		NameEN:        req.NameEN,
		NameTH:        req.NameTH,
		AddressTH:     req.AddressTH,
		AddressEN:     req.AddressEN,
		Phone:         req.Phone,
		Email:         req.Email,
		Facebook:      req.Facebook,
		X:             req.X,
		Instagram:     req.Instagram,
		DescriptionEN: req.DescriptionEN,
		DescriptionTH: req.DescriptionTH,
		ImageURL:      req.ImageURL,
		OpeningTime:   openingTime, // แปลงเป็น time.Time
		ClosingTime:   closingTime, // แปลงเป็น time.Time
	}

	// ใช้ tx เพื่อทำการสร้าง Cafe ภายใน transaction
	if err := tx.Create(&cafe).Error; err != nil {
		tx.Rollback() // ย้อนกลับหากมีข้อผิดพลาด
		return out, err
	}

	// ถ้าไม่มีข้อผิดพลาด ให้ยืนยันการทำธุรกรรม
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // ย้อนกลับหาก Commit ล้มเหลว
		return out, err
	}

	// คืนค่าข้อมูล Cafe ที่ถูกสร้างแล้ว
	return cafe, nil
}

func (s *CafeService) UpdateCafe(id string, req dto.CafeRequest) (models.Cafes, error) {
	var cafe models.Cafes
	tx := s.db.Begin()

	// ค้นหาข้อมูล Cafe ด้วย ID
	if err := s.db.Where("id = ?", id).First(&cafe).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return cafe, errors.New("Cafe not found")
		}
		return cafe, err
	}

	// พิมพ์ค่าของ req เพื่อตรวจสอบข้อมูลที่ส่งมา
	fmt.Printf("Request to Update Cafe: %+v\n", req)

	// สร้าง map สำหรับอัปเดต
	updateData := map[string]interface{}{}

	// อัปเดตเฉพาะฟิลด์ที่ไม่เป็นค่าเริ่มต้น
	if req.NameEN != "" {
		updateData["name_en"] = req.NameEN
	}
	if req.NameTH != "" {
		updateData["name_th"] = req.NameTH
	}
	if req.AddressTH != "" {
		updateData["address_th"] = req.AddressTH
	}
	if req.AddressEN != "" {
		updateData["address_en"] = req.AddressEN
	}
	if req.Phone != "" {
		updateData["phone"] = req.Phone
	}
	if req.Email != "" {
		updateData["email"] = req.Email
	}
	if req.Facebook != "" {
		updateData["facebook"] = req.Facebook
	}
	if req.X != "" {
		updateData["x"] = req.X
	}
	if req.Instagram != "" {
		updateData["instagram"] = req.Instagram
	}
	if req.DescriptionEN != "" {
		updateData["description_en"] = req.DescriptionEN
	}
	if req.DescriptionTH != "" {
		updateData["description_th"] = req.DescriptionTH
	}
	if req.ImageURL != "" {
		updateData["image_url"] = req.ImageURL
	}

	// แปลง OpeningTime และ ClosingTime จาก string เป็น time.Time (หากมีการอัปเดต)
	if req.OpeningTime != "" {
		openingTime, err := time.Parse("15:04:05", req.OpeningTime)
		if err != nil {
			tx.Rollback()
			return cafe, fmt.Errorf("invalid opening time format: %v", err)
		}
		updateData["opening_time"] = openingTime
	}
	if req.ClosingTime != "" {
		closingTime, err := time.Parse("15:04:05", req.ClosingTime)
		if err != nil {
			tx.Rollback()
			return cafe, fmt.Errorf("invalid closing time format: %v", err)
		}
		updateData["closing_time"] = closingTime
	}

	// ตรวจสอบว่า updateData มีข้อมูลก่อนอัปเดต
	if len(updateData) == 0 {
		tx.Rollback()
		return cafe, errors.New("no data to update")
	}

	// อัปเดต
	if err := s.db.Model(&cafe).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return cafe, err
	}

	// ถ้าไม่มีข้อผิดพลาด ให้ยืนยันการทำธุรกรรม
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // ย้อนกลับหาก Commit ล้มเหลว
		return cafe, err
	}
	// คืนค่าข้อมูล Cafe ที่อัปเดตแล้ว
	return cafe, nil
}

func (s *CafeService) GetCafeById(id string) (models.Cafes, error) {
	var cafe models.Cafes

	if err := s.db.Where("id = ?", id).First(&cafe).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return cafe, errors.New("Cafe not found")
		}
		return cafe, err
	}
	return cafe, nil
}

func (s *CafeService) GetAllCafe() ([]models.Cafes, error) {
	var out []models.Cafes
	// ใช้ Find เพื่อดึงข้อมูลทั้งหมดจากตาราง Cafes
	if err := s.db.Select("*").Find(&out).Error; err != nil {
		return nil, errors.New("Cafe Not Found")
	}
	// คืนค่าข้อมูลทั้งหมดที่ดึงมา
	return out, nil
}
