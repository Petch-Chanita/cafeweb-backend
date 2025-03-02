package models

import (
	"time"
)

// Cafe struct ใช้ BaseModel เพื่อให้มี fields ที่เราต้องการ
type Cafes struct {
	BaseModel               // Embed BaseModel
	NameEN        string    `gorm:"type:varchar(255);not null" json:"name_en"`
	NameTH        string    `gorm:"type:varchar(255);not null" json:"name_th"`
	AddressTH     string    `gorm:"type:varchar(255);not null" json:"address_th"`
	AddressEN     string    `gorm:"type:varchar(255);not null" json:"address_en"`
	Phone         string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Email         string    `gorm:"type:varchar(100)" json:"email,omitempty"`
	Facebook      string    `gorm:"type:varchar(100)" json:"facebook,omitempty"`
	X             string    `gorm:"type:varchar(100)" json:"x,omitempty"`
	Instagram     string    `gorm:"type:varchar(100)" json:"instagram,omitempty"`
	DescriptionEN string    `gorm:"type:text" json:"description_en,omitempty"`
	DescriptionTH string    `gorm:"type:text" json:"description_th,omitempty"`
	ImageURL      string    `gorm:"type:text" json:"image_url,omitempty"`
	OpeningTime   time.Time `gorm:"type:time;not null" json:"opening_time"`
	ClosingTime   time.Time `gorm:"type:time;not null" json:"closing_time"`
}
