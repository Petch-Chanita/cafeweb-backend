package models

type Images struct {
	BaseModel

	URL      string `gorm:"type:text;not null" json:"url"`      // URL ของรูปภาพ
	Filename string `gorm:"type:text;not null" json:"filename"` // ชื่อไฟล์ของรูปภาพ

	CafeID string `gorm:"type:uuid;index" json:"cafe_id"` // foreign key อ้างอิงจาก Cafe
	Cafe   Cafes  `gorm:"foreignKey:CafeID" json:"cafe"`  // ความสัมพันธ์กับ Cafe
}
