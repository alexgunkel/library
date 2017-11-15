package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
	"github.com/alexgunkel/library/repository"
	"strconv"
)

func CreateBook(c *gin.Context)  {
	var book models.Book
	c.Bind(&book)
	repository.AddBook(&book)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book item created successfully!", "data": book})
}

func ListBooks(c *gin.Context)  {
	books, err := repository.GetBooks()
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": books})
}

func GetBook(c *gin.Context)  {
	bookid := getId(c)
	book, err := repository.GetBookById(bookid)

	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &book})
}

func GetBookAuthors(c *gin.Context)  {
	bookid := getId(c)
	authors, err := repository.GetBookAuthors(bookid)

	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &authors})
}

func AddBookAuthor(c *gin.Context) {
	var author models.Author
	c.Bind(&author)
	bookid := getId(c)

	err := repository.AddBookAuthor(bookid, &author)

	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
}

func UpdateBook(c *gin.Context)  {
	bookid := getId(c)
	var book models.Book
	c.Bind(&book)
	repository.UpdateBook(bookid, &book)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully updated"})
}

func DeleteBook(c *gin.Context)  {
	bookid := getId(c)
	repository.DeleteBook(bookid)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &book})
}

func getId(c *gin.Context) int64 {
	bookid, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	return bookid
}
