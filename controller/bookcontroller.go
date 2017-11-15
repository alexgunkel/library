package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
)

func CreateBook(c *gin.Context)  {
	book := models.Book{Title: c.PostForm("title"), Subtitle: c.PostForm("subtitle"), Identifier: c.PostForm("identifier")}
	book.Authors = []models.Author{
		{
			FirstName: "Peter",
			LastName: "Müller",
		},
		{
			FirstName: "Hans",
			LastName: "Wurst",
		},
	}
	db.Save(&book)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book item created successfully!", "data": book})
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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &book})
}

func UpdateBook(c *gin.Context)  {
	var book models.Book
	bookid := c.Param("id")
	db.First(&book, bookid)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Book not found"})
		return
	}

	db.Model(&book).Update("title", c.PostForm("title"))
	db.Model(&book).Update("subtitle", c.PostForm("subtitle"))
	db.Model(&book).Update("identifier", c.PostForm("identifier"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully updated"})
}

func DeleteBook(c *gin.Context)  {
	var book models.Book
	bookid := c.Param("id")
	db.First(&book, bookid)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Could not find book"})
		return
	}

	db.Model(&book).Delete(&book, bookid)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &book})
}
