package routes

import (
	"cafeweb-backend/controllers"
	"cafeweb-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetUpRoutes - ฟังก์ชันสำหรับตั้งค่า routes
func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	// ====== Services ======
	userService := services.NewUserService(db)
	cafeService := services.NewCafeService(db)
	uploadService := services.NewUploadService(db)
	aboutService := services.NewAboutService(db)
	productService := services.NewProductService(db)

	// ====== Controllers ======
	userController := controllers.NewUserController(userService)
	cafeController := controllers.NewCafeController(cafeService)
	uploadController := controllers.NewUploadController(uploadService)
	aboutController := controllers.NewAboutController(aboutService)
	productController := controllers.NewProductController(productService, uploadService)

	// ====== Static Files ======
	router.Static("/uploads", "./uploads")

	// ====== Routes ======

	// Upload routes
	router.POST("/api/upload/:cafe_id", uploadController.UploadImage)

	// Auth routes
	authApi := router.Group("/api/auth")
	{
		authApi.POST("/login-admin", userController.Login)
		authApi.POST("/register", userController.RegisterUser)
	}

	// User routes
	userApi := router.Group("/api/users")
	{
		userApi.GET("/", userController.GetAllUsers)
		userApi.GET("/:id", userController.GetUserById)
		userApi.PUT("/:id", userController.UpdateUser)
		userApi.DELETE("/", userController.DeleteUser)
	}

	// Cafe routes
	cafeApi := router.Group("/api/cafe")
	{
		cafeApi.GET("/", cafeController.GetAllCafe)
		cafeApi.GET("/:id", cafeController.GetCafeById)
		cafeApi.PUT("/update-cafe/:id", cafeController.UpdateCafe)
		cafeApi.POST("/create-cafe", cafeController.CreateCafe)
	}

	// About routes
	aboutApi := router.Group("/api/abouts")
	{
		aboutApi.GET("/:cafe_id", aboutController.GetAboutByCafeID)
		aboutApi.POST("/", aboutController.CreateAboutHandler)
		aboutApi.PUT("/:id", aboutController.UpdateAboutHandler)
	}

	//Product routes
	productApi := router.Group("/api/product")
	{
		productApi.POST("/", productController.CreateProduct)
		productApi.POST("/:id", productController.UpdateProduct)
		productApi.GET("/", productController.GetProduct)
		productApi.GET("/:id", productController.GetProductByID)
		productApi.DELETE("/", productController.DeleteMultipleProducts)

		productApi.POST("/category", productController.CreateCategories)
		productApi.GET("/category", productController.GetCategories)
	}
}
