package dto

import "cafeweb-backend/models"

type ProductRequest struct {
	ProductName string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required"`
	ImageURL    string `json:"image_url"`
	Status      string `json:"status" binding:"required"`
	CategoryID  string `json:"category_id" binding:"required"`
	CafeID      string `json:"cafe_id" binding:"required"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ID           string            `json:"id"`
	ProductName  string            `json:"name"`
	Price        string            `json:"price"`
	ImageURL     string            `json:"image_url"`
	CategoryName string            `json:"type"`
	CategoryID   string            `json:"category_id,omitempty"`
	CafeID       string            `json:"cafe_id"`
	Status       models.StatusType `json:"status"`
	Description  string            `json:"description,omitempty"`
}

type CategoryRequest struct {
	CafeID       string `json:"cafe_id" binding:"required"`
	CategoryName string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID           string ` json:"id"`
	CafeID       string ` json:"cafe_id"`
	CategoryName string `json:"name"`
}
