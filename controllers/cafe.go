package controllers

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CafeController - ฟังก์ชันที่ใช้ในการรับคำขอจาก client และใช้ service ในการจัดการ
type CafeController struct {
	CafeService *services.CafeService
}

// NewCafeController - ฟังก์ชันสำหรับสร้าง CafeController
func NewCafeController(service *services.CafeService) *CafeController {
	return &CafeController{
		CafeService: service,
	}
}

func (c *CafeController) CreateCafe(ctx *gin.Context) {
	var cafe dto.CafeRequest

	if err := ctx.ShouldBindJSON(&cafe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.CafeService.CreateCafe(cafe)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *CafeController) UpdateCafe(ctx *gin.Context) {
	id := ctx.Param("id")
	var cafe dto.CafeRequest

	if err := ctx.ShouldBindJSON(&cafe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.CafeService.UpdateCafe(id, cafe)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *CafeController) GetCafeById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.CafeService.GetCafeById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *CafeController) GetAllCafe(ctx *gin.Context) {

	data, err := c.CafeService.GetAllCafe()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
