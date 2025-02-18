package models

// About struct สำหรับสร้างตาราง abouts
type Abouts struct {
	BaseModel // Embed BaseModel เพื่อให้มี fields เช่น ID, CreatedAt, UpdatedAt, DeletedAt

	CafeID string `gorm:"type:uuid;index" json:"cafe_id"` // Foreign key ที่อ้างอิงจาก Cafe
	Cafe   Cafes  `gorm:"foreignKey:CafeID" json:"cafe"`  // ความสัมพันธ์กับ Cafe

	AboutEn *string `gorm:"type:text" json:"about_en"` // คำอธิบายภาษาอังกฤษ
	AboutTh *string `gorm:"type:text" json:"about_th"` // คำอธิบายภาษาไทย
}
