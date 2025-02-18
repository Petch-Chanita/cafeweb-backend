package routes

import (
	"cafeweb-backend/config"
	"cafeweb-backend/controllers"
	"cafeweb-backend/services"

	"github.com/gin-gonic/gin"
)

// SetUpRoutes - ฟังก์ชันสำหรับตั้งค่า routes
func SetUpRoutes(router *gin.Engine) {
	// สร้าง instance ของ UserService และ UserController
	userService := services.NewUserService(config.DB)
	userController := controllers.NewUserController(userService)

	// สร้าง instance ของ CafeService และ UserController
	cafeService := services.NewCafeService(config.DB)
	cafeController := controllers.NewCafeController(cafeService)

	uploadService := services.NewUploadService(config.DB)
	uploadController := controllers.NewUploadController(uploadService)

	router.Static("/uploads", "./uploads")
	// ตั้งค่า route สำหรับการอัปโหลด
	router.POST("/api/upload/:cafe_id", uploadController.UploadImage)

	authApi := router.Group("/api/auth")
	{
		authApi.POST("/login-admin", userController.Login)
		authApi.POST("/register", userController.RegisterUser)
	}
	userApi := router.Group("/api/users")
	{
		userApi.GET("/", userController.GetAllUsers)
		userApi.GET("/:id", userController.GetUserById)
		userApi.PUT("/:id", userController.UpdateUser)
		userApi.DELETE("/", userController.DeleteUser)
	}
	cafeApi := router.Group("/api/cafe")
	{
		cafeApi.GET("/", cafeController.GetAllCafe)
		cafeApi.GET("/:id", cafeController.GetCafeById)
		cafeApi.PUT("/update-cafe/:id", cafeController.UpdateCafe)
		cafeApi.POST("/create-cafe", cafeController.CreateCafe)
	}
}
