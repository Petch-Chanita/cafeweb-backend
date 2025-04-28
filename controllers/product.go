package controllers

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/services"
	"cafeweb-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductController - struct ควบคุมเกี่ยวกับ Product
type ProductController struct {
	ProductService *services.ProductService
}

// NewProductController - ฟังก์ชันสร้าง ProductController
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

// CreateProduct - API สำหรับสร้าง Product
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product dto.ProductRequest

	_, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.ProductService.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
		"data":    data,
	})
}

func (c *ProductController) GetProduct(ctx *gin.Context) {

	claims, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cafeID, ok := claims["cafe_id"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "cafe_id not found in token"})
		return
	}

	data, err := c.ProductService.GetProduct(cafeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

// CreateCategories - API สำหรับสร้าง Category
func (c *ProductController) CreateCategories(ctx *gin.Context) {
	var category dto.CategoryRequest
	_, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ProductService.CreateCategories(category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "category created successfully",
	})
}
