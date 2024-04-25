package app

import (
	"go/api/internal/domain"
	entities "go/api/internal/entities"
	middleware "go/api/pkg/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(router *gin.Engine, bookUC domain.BookUseCase) {

	// Call Middleware Before the Handler
	router.Use(middleware.TestMiddleware())

	router.GET("/books", func(c *gin.Context) {
		books, err := bookUC.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	router.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}
		book, err := bookUC.GetBookByID(uint(bookID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, book)
	})

	router.POST("/books", func(c *gin.Context) {
		var book entities.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := bookUC.CreateBook(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
	})

	router.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}
		var book entities.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := bookUC.UpdateBook(uint(bookID), &book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	})

	router.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}
		if err := bookUC.DeleteBook(uint(bookID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})
}
