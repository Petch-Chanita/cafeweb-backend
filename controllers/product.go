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
	UploadService  *services.UploadService
}

// NewProductController - ฟังก์ชันสร้าง ProductController
func NewProductController(productService *services.ProductService, uploadService *services.UploadService) *ProductController {
	return &ProductController{
		ProductService: productService,
		UploadService:  uploadService,
	}
}

// CreateProduct - API สำหรับสร้าง Product
func (c *ProductController) CreateProduct(ctx *gin.Context) {

	token, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID, ok := token["id"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	categoryID := ctx.PostForm("category_id")
	status := ctx.PostForm("status")
	description := ctx.PostForm("description")
	cafeID := ctx.PostForm("cafe_id")

	file, err := ctx.FormFile("image")

	if err != nil || file == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	imageURL, err := c.UploadService.UploadFile(cafeID, userID, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req := dto.ProductRequest{
		ProductName: name,
		Price:       price,
		CategoryID:  categoryID,
		Status:      status,
		Description: description,
		CafeID:      cafeID,
		ImageURL:    imageURL,
	}

	data, err := c.ProductService.CreateProduct(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
		"data":    data,
	})
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	token, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID, ok := token["id"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	categoryID := ctx.PostForm("category_id")
	status := ctx.PostForm("status")
	description := ctx.PostForm("description")
	cafeID := ctx.PostForm("cafe_id")

	var imageURL string
	file, err := ctx.FormFile("image")
	if err == nil && file != nil {
		imageURL, err = c.UploadService.UploadFile(cafeID, userID, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		imageURL = ctx.PostForm("image_url")
	}

	req := dto.ProductRequest{
		ProductName: name,
		Price:       price,
		CategoryID:  categoryID,
		Status:      status,
		Description: description,
		CafeID:      cafeID,
		ImageURL:    imageURL,
	}

	data, err := c.ProductService.UpdateProduct(cafeID, productID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product updated successfully",
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

func (c *ProductController) GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")

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

	data, err := c.ProductService.GetProductByID(cafeID, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

func (c *ProductController) DeleteMultipleProducts(ctx *gin.Context) {
	var request struct {
		IDs []string `json:"ids"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.ProductService.DeleteMultipleProducts(request.IDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Products deleted successfully"})
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

func (c *ProductController) GetCategories(ctx *gin.Context) {
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

	data, err := c.ProductService.GetCategories(cafeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, data)
}
