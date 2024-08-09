package controllers

import (
	"fmt"
	"net/http"

	"github.com/Sattuu21/go-bookstore/pkg/models"
	"github.com/Sattuu21/go-bookstore/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the books as JSON
	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId") //This line retrieves the bookId parameter from the URL. For example, if the URL is /book/123, bookId would be "123".
	var book models.Book
	if err := models.DB.First(&book, bookId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Ensure the ID is not manually set or passed by the frontend
	if book.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should not be set manually"})
		return
	}

	// The database will automatically generate the ID
	if err := models.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.Book
	if err := models.DB.First(&book, bookId).Error; err != nil { //first checking if the book exists
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Delete(&book, bookId).Error; err != nil { // using unscoped to hard delete , as gorm suports soft vs hard delete, but will not use and store the record in db
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Book deleted with id: %s\n", bookId)

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.Book

	// Fetch the book by ID correctly
	if err := models.DB.First(&book, "id = ?", bookId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedBook models.Book
	if err := utils.ParseBody(c.Request, &updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing"})
		return
	}

	book.Name = updatedBook.Name // update the fields
	book.Author = updatedBook.Author
	book.Type = updatedBook.Type

	// Save the updated book
	if err := models.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}
