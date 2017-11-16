package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
	"github.com/alexgunkel/library/repository"
)

func CreateAuthor(c *gin.Context)  {
	var author models.Author
	c.Bind(&author)
	repository.AddAuthor(&author)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Author item created successfully!", "data": author})
}

func ListAuthors(c *gin.Context)  {
	authors, err := repository.GetAuthors()
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": authors})
}

func GetAuthor(c *gin.Context)  {
	authorid := getId(c)
	author, err := repository.GetAuthorById(authorid)

	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &author})
}

func UpdateAuthor(c *gin.Context)  {
	var author models.Author
	authorid := getId(c)
	c.Bind(&author)
	repository.UpdateAuthor(authorid, &author)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully updated"})
}

func DeleteAuthor(c *gin.Context)  {
	err := repository.DeleteAuthor(getId(c))
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
