package services

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"cafeweb-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UploadService struct {
	db *gorm.DB
}

func NewUploadService(db *gorm.DB) *UploadService {
	return &UploadService{db: db}
}

// ฟังก์ชันสำหรับอัปโหลดไฟล์
func (s *UploadService) UploadFile(cafeID string, userID string, file *multipart.FileHeader) (string, error) {
	if s.db == nil {
		log.Println("❌ UploadService.db is nil")
		return "", fmt.Errorf("internal server error: database not initialized")
	}

	tx := s.db.Begin()
	if file == nil {
		return "", fmt.Errorf("file header is nil")
	}

	log.Println("File name:", file.Filename)

	uploadDir := "uploads"

	// ตรวจสอบว่ามีโฟลเดอร์ `uploads` หรือไม่ ถ้าไม่มีให้สร้างใหม่
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create upload directory: %v", err)
		}
	}
	// กำหนดที่เก็บไฟล์
	uploadPath := "uploads/" + file.Filename

	log.Println("File name::", file.Filename)
	log.Println("uploadPath:", uploadPath)

	// เปิดไฟล์สำหรับเขียน
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// สร้างไฟล์ใหม่
	dst, err := os.Create(uploadPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// คัดลอกไฟล์จาก src ไปยัง dst
	_, err = dst.ReadFrom(src)
	if err != nil {
		return "", err
	}

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080" // ค่าเริ่มต้นถ้าไม่ได้ตั้งค่า BASE_URL
	}

	// สร้าง URL ของไฟล์ที่อัปโหลด
	fileUrl := fmt.Sprintf("%s/uploads/%s", baseUrl, file.Filename)

	// บันทึก URL และชื่อไฟล์ลงในฐานข้อมูล
	image := models.Images{
		URL:      fileUrl,
		Filename: file.Filename,
		CafeID:   cafeID,
		UserID:   userID,
	}

	if image.URL == "" || image.Filename == "" {
		return "", fmt.Errorf("invalid image data")
	}
	log.Println("image", &image)

	if err := tx.Create(&image).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to create image record: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", err
	}

	return fileUrl, nil
}
