package main

import (
	"fmt"
	co "go/api/configs"
	"go/api/internal/app"
	middleware "go/api/internal/middlewares"
	"go/api/internal/repository"
	"go/api/internal/usecase"
	"go/api/migrations"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load.env file:", err)
	}
	// Initialize Gin router
	router := gin.Default()

	// Initialize Gorm database connection
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",
		co.DBConfig().DBHost, co.DBConfig().DBUser, co.DBConfig().DBName, co.DBConfig().DBPort, co.DBConfig().DBPassword)

	// Open database connection
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Run database migrations
	if err := migrations.MigrateRun(db); err != nil {
		fmt.Println("Failed to run migrations:", err)
		return
	}

	// CORS
	router.Use(cors.Default())

	// Save Log
	router.Use(middleware.SaveLog)

	// Initialize repositories
	bookRepo := repository.NewBookRepository(db)

	// Initialize use cases
	bookUC := usecase.NewBookUseCase(bookRepo)

	// Setup routes
	app.SetupBookRoutes(router, bookUC)

	port := os.Getenv("PORT")
	// Start server
	router.Run(":" + port)

}
