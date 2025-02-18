package models

// Category struct สำหรับสร้างตาราง categories
type Categories struct {
	BaseModel // รวมฟิลด์ที่เป็นพื้นฐาน เช่น ID, CreatedAt, UpdatedAt, DeletedAt

	CafeID       string `gorm:"type:uuid;index" json:"cafe_id"`         // foreign key อ้างอิงจาก Cafe
	CategoryName string `gorm:"type:varchar(100);not null" json:"name"` // ชื่อหมวดหมู่
	Cafe         Cafes  `gorm:"foreignKey:CafeID" json:"cafe"`          // ความสัมพันธ์กับ Cafe
}
