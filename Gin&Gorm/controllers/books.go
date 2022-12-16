package controllers

import (
	"github.com/LitxGoGo/ginAndGorm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"Author" binding:"required"`
}

type UpdateBookInput struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"Author"`
}

// FindBooks Get /books
//Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// CreateBook POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBook GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//fmt.Println(book)
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//models.DB.Model(&book).Updates(UpdateBookInput{
	//	Title:  "123133",
	//	Author: "4684984",
	//})
	//在全结构体更新的时候所有的字段要一一对应
	models.DB.Model(&book).Updates(input)

	//models.DB.Model(&book).Update("title", input.Title)
	//models.DB.Model(&book).Update("Author", input.Author)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&book).Delete(&book)

	c.JSON(200, gin.H{
		"data": true,
	})
}
