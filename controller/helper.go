package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getId(c *gin.Context) int64 {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	return id
}

func ok(c *gin.Context, obj interface{})  {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": obj})
}