package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

	// ตรวจสอบว่า token เป็น valid และไม่หมดอายุ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// ตรวจสอบว่า token หมดอายุหรือไม่
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token has expired")
			}
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// ตรวจสอบและดึงข้อมูลจาก JWT Token
func GetClaimsFromToken(ctx *gin.Context) (jwt.MapClaims, error) {
	// รับ Authorization header จาก request
	authHeader := ctx.GetHeader("Authorization")

	// ตรวจสอบว่า header มีหรือไม่
	if authHeader == "" {
		return nil, fmt.Errorf("Authorization header is required")
	}

	// แยก Bearer Token
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// ตรวจสอบความถูกต้องของ Token
	claims, err := ParseToken(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid or expired token: %v", err)
	}

	return claims, nil
}
