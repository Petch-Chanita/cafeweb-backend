package controllers

import (
	"net/http"

	"cafeweb-backend/models"
	"cafeweb-backend/services"
	"cafeweb-backend/utils"

	"github.com/gin-gonic/gin"
)

// AboutController - struct ควบคุมเกี่ยวกับ Abouts
type AboutController struct {
	AboutService *services.AboutService
}

// NewAboutController - ฟังก์ชันสร้าง AboutController
func NewAboutController(service *services.AboutService) *AboutController {
	return &AboutController{AboutService: service}
}

// CreateAboutHandler - API สำหรับสร้าง About
func (c *AboutController) CreateAboutHandler(ctx *gin.Context) {
	var about models.Abouts

	_, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.AboutService.CreateAbout(&about); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, about)
}

// UpdateAboutHandler - API สำหรับอัปเดต About
func (c *AboutController) UpdateAboutHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedData models.Abouts

	_, err := utils.GetClaimsFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.AboutService.UpdateAbout(id, &updatedData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "About updated successfully"})
}

func (c *AboutController) GetAboutByCafeID(ctx *gin.Context) {
	cafeID := ctx.Param("cafe_id")

	// เรียกใช้ Service เพื่อดึงข้อมูล About
	about, err := c.AboutService.GetAboutByCafeID(cafeID)
	if err != nil {
		// ถ้าเกิด error ในการดึงข้อมูล
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if about == nil {
		// ถ้าไม่พบข้อมูล
		ctx.JSON(http.StatusNotFound, gin.H{"message": "About not found for this cafe"})
		return
	}

	// ส่งข้อมูลที่ได้กลับไปใน response
	ctx.JSON(http.StatusOK, about)
}
