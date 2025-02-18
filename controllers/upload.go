package controllers

import (
	"cafeweb-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	UploadService *services.UploadService
}

func NewUploadController(uploadService *services.UploadService) *UploadController {
	return &UploadController{UploadService: uploadService}
}

// ฟังก์ชันสำหรับอัปโหลดไฟล์
func (c *UploadController) UploadImage(ctx *gin.Context) {
	cafeID := ctx.Param("cafe_id")

	// รับไฟล์จาก request
	_, FileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
		return
	}

	// เรียกใช้ Service เพื่ออัปโหลดไฟล์
	fileUrl, err := c.UploadService.UploadFile(cafeID, FileHeader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่ง response กลับไป
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded & linked to cafe",
		"url":     fileUrl,
	})
}
