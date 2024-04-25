package main

import (
	"fmt"
	co "go/api/configs"
	"go/api/internal/app"
	"go/api/internal/repository"
	"go/api/internal/usecase"
	"go/api/pkg/middlewares"
	"go/api/pkg/migrations"
	"go/api/pkg/tasks"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/robfig/cron/v3"
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

	// Open database gorm connection
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: newLogger,
	})

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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(corsConfig))

	// Save Log
	router.Use(middlewares.SaveLog)

	// Task Scheduler
	c := cron.New()
	c.AddFunc("@every 1s", tasks.HelloTask)
	c.Start()

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
