package dto

type CafeRequest struct {
	NameEN        string `json:"name_en" binding:"required"`
	NameTH        string `json:"name_th" binding:"required"`
	AddressTH     string `json:"address_th"`
	AddressEN     string `json:"address_en"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Facebook      string `json:"facebook"`
	X             string `json:"x"`
	Instagram     string `json:"instagram"`
	DescriptionEN string `json:"description_en"`
	DescriptionTH string `json:"description_th"`
	ImageURL      string `json:"image_url"`
	OpeningTime   string `json:"opening_time" binding:"required"`
	ClosingTime   string `json:"closing_time" binding:"required"`
}
