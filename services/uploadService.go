package services

import (
	"fmt"
	"mime/multipart"
	"os"

	"cafeweb-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UploadService struct {
	DB *gorm.DB
}

func NewUploadService(db *gorm.DB) *UploadService {
	return &UploadService{DB: db}
}

// ฟังก์ชันสำหรับอัปโหลดไฟล์
func (s *UploadService) UploadFile(cafeID string, file *multipart.FileHeader) (string, error) {
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
	}

	if err := s.DB.Create(&image).Error; err != nil {
		return "", err
	}

	return fileUrl, nil
}
