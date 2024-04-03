package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveLog(c *gin.Context) {
	// Generate a request ID
	requestID := uuid.New().String()

	// Create a new log file every day
	logFileName := fmt.Sprintf("./logs/%s.log", time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return // Stop processing if log file creation fails
	}
	defer f.Close()

	// Log to file
	log.SetOutput(f)
	start := time.Now()

	// Log request details with IP address and request ID
	log.Printf("ReqID: %s | Method: %s %s | IP: %s ", requestID, c.Request.Method, c.Request.URL.Path, c.ClientIP())
	c.Next()

	// Log response status and latency
	end := time.Now()
	latency := end.Sub(start)
	log.Printf("ReqID: %s | Method: %s %s | ResponseStatus: %d | Latency: %s ", requestID, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency.String())
}

// AuthMiddleware is an example middleware for authentication
func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Middleware logic to check authentication
		// Example: Check if the user is logged in
		// If not authenticated, you can return an error or redirect
		// For simplicity, this example middleware allows all requests
		fmt.Println("1111111111111111111111111111111111111111")
		c.Next() // Call the next handler
	}
}
