package controllers

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController - ฟังก์ชันที่ใช้ในการรับคำขอจาก client และใช้ service ในการจัดการ
type UserController struct {
	UserService *services.UserService
}

// NewUserController - ฟังก์ชันสำหรับสร้าง UserController
func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// Login - ฟังก์ชันที่ใช้จัดการการ login
func (c *UserController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียกใช้ service เพื่อทำการ login
	token, err := c.UserService.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var input dto.UserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Username == "" || input.Password == "" || input.CafeID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	user, err := c.UserService.AddUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var input dto.UserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedUser, err := c.UserService.UpdateUser(id, input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var req dto.RequestData

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deletedUser, err := c.UserService.DeleteUser(req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deletedUser)
}
