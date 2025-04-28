package models

// Product struct สำหรับสร้างตาราง products
type Products struct {
	BaseModel // รวมฟิลด์ที่เป็นพื้นฐาน เช่น ID, CreatedAt, UpdatedAt, DeletedAt

	ProductName string `gorm:"type:varchar(255);not null" json:"productName"` // ชื่อสินค้า
	Price       string `gorm:"type:decimal(10,2);not null" json:"price"`      // ราคา
	CafeID      string `gorm:"type:uuid;index" json:"cafe_id"`                // foreign key อ้างอิงจาก Cafe
	CategoryID  string `gorm:"type:uuid;index" json:"category_id"`            // foreign key อ้างอิงจาก Category

	ImageURL    string `gorm:"type:text" json:"image_url"`   // URL รูปภาพ
	Description string `gorm:"type:text" json:"description"` // คำอธิบายสินค้า

	Cafe     Cafes      `gorm:"foreignKey:CafeID" json:"cafe"`         // ความสัมพันธ์กับ Cafe
	Category Categories `gorm:"foreignKey:CategoryID" json:"category"` // ความสัมพันธ์กับ Category
}
