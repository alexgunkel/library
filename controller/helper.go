package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func getId(c *gin.Context) int64 {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	return id
}
