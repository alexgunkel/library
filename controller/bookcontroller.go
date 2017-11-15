package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
)

func CreateBook(c *gin.Context)  {
	book := models.Book{Title: c.PostForm("title"), Author: c.PostForm("author"), Subtitle: c.PostForm("subtitle"), Identifier: c.PostForm("identifier")}
	db.Save(&book)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book item created successfully!", "resourceId": book.ID})
}

func ListBooks(c *gin.Context)  {
	var books []models.Book

	db.Find(&books)

	if len(books) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No books found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": books})
}

func GetBook(c *gin.Context)  {
	var book models.Book
	bookid := c.Param("id")
	db.First(&book, bookid)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Could not find book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": book})
}
