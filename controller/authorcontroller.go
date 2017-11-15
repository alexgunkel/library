package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alexgunkel/library/models"
)

func CreateAuthor(c *gin.Context)  {
	author := models.Author{FirstName: c.PostForm("first_name"), LastName: c.PostForm("last_name")}
	db.Save(&author)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Author item created successfully!", "data": author})
}

func ListAuthors(c *gin.Context)  {
	var authors []models.Author

	db.Find(&authors)

	if len(authors) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No authors found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": authors})
}

func GetAuthor(c *gin.Context)  {
	var author models.Author
	authorid := c.Param("id")
	db.First(&author, authorid)

	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Could not find author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &author})
}

func UpdateAuthor(c *gin.Context)  {
	var author models.Author
	authorid := c.Param("id")
	db.First(&author, authorid)

	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Author not found"})
		return
	}

	db.Model(&author).Update("first_name", c.PostForm("first_name"))
	db.Model(&author).Update("last_name", c.PostForm("last_name"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully updated"})
}

func DeleteAuthor(c *gin.Context)  {
	var author models.Author
	authorid := c.Param("id")
	db.First(&author, authorid)

	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Could not find author"})
		return
	}

	db.Model(&author).Delete(&author, authorid)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &author})
}
