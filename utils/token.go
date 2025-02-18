package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var secretKey = "JWT_SECRET" // คีย์สำหรับการสร้าง JWT

// GenerateToken - ฟังก์ชันสร้าง JWT Token
func GenerateToken(userID uuid.UUID, userName string, role string, image *string) (string, error) {
	claims := jwt.MapClaims{
		"id":       userID.String(),                       // รหัสผู้ใช้
		"iat":      time.Now().Unix(),                     // เวลาที่ Token ถูกสร้าง
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // หมดอายุใน 24 ชั่วโมง
		"username": userName,                              // ชื่อผู้ใช้
		"role":     role,
		"image":    image,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken - ฟังก์ชันสำหรับตรวจสอบและดึงข้อมูลจาก JWT
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	// แยกข้อมูลจาก JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่าเป็นการใช้ Signing Method ที่ถูกต้อง
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่า token เป็น valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
