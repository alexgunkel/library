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
		vi.GET("/:id/authors", controller.GetBookAuthors)
		vi.POST("/:id/authors", controller.AddBookAuthor)
		vi.POST("/", controller.CreateBook)
		vi.PUT("/:id", controller.UpdateBook)
		vi.DELETE("/:id", controller.DeleteBook)
	}
	author := router.Group("api/v1/authors")
	{
		author.GET("/", controller.ListAuthors)
		author.GET("/:id", controller.GetAuthor)
		author.POST("/", controller.CreateAuthor)
		author.PUT("/:id", controller.UpdateAuthor)
		author.DELETE("/:id", controller.DeleteAuthor)
	}
	router.Run()
}
