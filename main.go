package main

import (
	"github.com/gin-gonic/gin"
	"github.com/alexgunkel/library/controller"
)


func main()  {
	router := gin.Default()

	vi := router.Group("api/v1/books")
	{
		vi.GET("/", controller.ListBooks)
		vi.GET("/:id", controller.GetBook)
		vi.POST("/", controller.CreateBook)
		//vi.PUT("/:id", updateBook)
		//vi.DELETE("/:id", deleteBook)
	}
	router.Run()
}
