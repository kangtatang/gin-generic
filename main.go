package main

import (
	"go-gin-generic/handlers"
	"go-gin-generic/models"
	"go-gin-generic/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Konfigurasi koneksi PostgreSQL
	dsn := "host=localhost user=postgres password=yourpassword dbname=gin_generic port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{})

	// Inisialisasi repository dan handler
	userRepo := &repositories.GenericRepository[models.User]{DB: db}
	userHandler := &handlers.UserHandler{Repo: userRepo}

	// Inisialisasi Gin
	r := gin.Default()

	// Routing
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)

	// Jalankan server
	r.Run(":8885")
}
