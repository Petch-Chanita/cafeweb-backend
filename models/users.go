package models

// User struct สำหรับสร้างตาราง users
type Users struct {
	BaseModel // Embed BaseModel เพื่อให้มี fields เช่น ID, CreatedAt, UpdatedAt, DeletedAt

	Username string  `gorm:"type:varchar(100);unique;not null" json:"username"` // ชื่อผู้ใช้
	Password string  `gorm:"type:varchar(100);not null" json:"password"`        // รหัสผ่าน
	Role     string  `gorm:"type:varchar(50);not null" json:"role"`             // บทบาท (admin, user, etc.)
	Image    *string `gorm:"type:text" json:"image,omitempty"`                  // URL รูปภาพของผู้ใช้

	CafeID string `gorm:"type:uuid;not null" json:"cafe_id"` // Foreign key ที่อ้างอิงจาก Cafe
	Cafe   Cafes  `gorm:"foreignKey:CafeID" json:"cafe"`     // ความสัมพันธ์กับ Cafe
}
