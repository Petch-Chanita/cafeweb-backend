package main

import (
	"fmt"
	"log"
	"os"

	"cafeweb-backend/config"
	"cafeweb-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// โหลด ENV และเชื่อมต่อฐานข้อมูล
	config.InitDB()

	r := gin.Default()

	// ตั้งค่า CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // อนุญาตให้ทุกโดเมนสามารถเข้าถึงได้
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // เมธอดที่อนุญาต
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // เฮดเดอร์ที่อนุญาต
		AllowCredentials: true,                                                // ถ้าต้องการส่งข้อมูล Cookie หรือ Credential
	}))

	// ตั้งค่า routes
	routes.SetUpRoutes(r)

	// เรียกใช้ฟังก์ชันเพื่อทำ migration
	// migration.RunMigration()

	jwtSecret := os.Getenv("JWT_SECRET")
	apiKey := os.Getenv("API_KEY")

	fmt.Println("🔑 JWT Secret:", jwtSecret)
	fmt.Println("🔑 API Key:", apiKey)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
