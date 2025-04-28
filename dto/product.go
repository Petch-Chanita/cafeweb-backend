package dto

type ProductRequest struct {
	ProductName string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required"`
	ImageURL    string `json:"image_url"`
	CategoryID  string `json:"category_id" binding:"required"`
	CafeID      string `json:"cafe_id" binding:"required"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ProductName string `json:"name"`
	Price       string `json:"price"`
	ImageURL    string `json:"image_url"`
	CategoryID  string `json:"category_id"`
	CafeID      string `json:"cafe_id"`
	Description string `json:"description,omitempty"`
}

type CategoryRequest struct {
	CafeID       string `json:"cafe_id" binding:"required"`
	CategoryName string `json:"name" binding:"required"`
}
