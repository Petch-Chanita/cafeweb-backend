package services

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/models"
	"cafeweb-backend/utils"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *gorm.DB
}

// NewUserService - ฟังก์ชันสำหรับสร้าง UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// LoginUser - ฟังก์ชันที่ใช้ตรวจสอบ username และ password
func (s *UserService) LoginUser(username, password string) (string, error) {
	var user models.Users
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("Invalid credentials")
	}

	// สร้าง JWT token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role, user.Image)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) AddUser(req dto.UserRequest) (models.Users, error) {
	tx := s.db.Begin()

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return models.Users{}, err // หยุดหากเกิดข้อผิดพลาดในการเข้ารหัสรหัสผ่าน
	}

	user := models.Users{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
		CafeID:   req.CafeID,
	}

	if err := s.db.Create(&user).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	// ใช้ Preload เพื่อดึงข้อมูล Cafe ที่เชื่อมโยงกับผู้ใช้
	if err := tx.Preload("Cafe").First(&user).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	// ถ้าไม่มีข้อผิดพลาด ให้ยืนยันการทำธุรกรรม
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // ย้อนกลับหาก Commit ล้มเหลว
		return user, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserById(id string) (*models.Users, error) {
	var user models.Users
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(id string, updatedData dto.UserRequest) (*models.Users, error) {
	var user models.Users
	if err := s.db.First(&user, "id = ? and cafe_id = ?", id, updatedData.CafeID).Error; err != nil {
		return nil, err
	}

	user.Username = updatedData.Username
	user.Password = updatedData.Password
	user.Role = updatedData.Role

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) DeleteUser(req dto.RequestData) (*models.Users, error) {
	var user models.Users
	if err := s.db.First(&user, "id = ? AND cafe_id = ?", req.ID, req.CafeID).Error; err != nil {
		return nil, err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
