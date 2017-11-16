package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
	"github.com/alexgunkel/library/repository"
)

func CreateAuthor(c *gin.Context)  {
	var author models.Author
	err := c.Bind(&author)
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	go repository.AddAuthor(&author)
	ok(c, &author)
}

func ListAuthors(c *gin.Context)  {
	authors, err := repository.GetAuthors()
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	ok(c, &authors)
}

func GetAuthor(c *gin.Context)  {
	authorid := getId(c)
	author, err := repository.GetAuthorById(authorid)

	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	ok(c, &author)
}

func UpdateAuthor(c *gin.Context)  {
	var author models.Author
	authorid := getId(c)
	c.Bind(&author)
	go repository.UpdateAuthor(authorid, &author)

	ok(c, &author)
}

func DeleteAuthor(c *gin.Context)  {
	err := repository.DeleteAuthor(getId(c))
	if nil != err {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
