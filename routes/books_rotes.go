package routes

import (
	"github.com/gin-gonic/gin"
	"kurdi-go/controllers"
	"kurdi-go/middlewares"
)

func BooksRoutes(booksRouter *gin.Engine) {
	booksRouter.Use(middlewares.AuthenticationMiddleware)
	booksRouter.GET("/get", func(context *gin.Context) {
		var a = 5
		var p = &a
		context.JSON(200, gin.H{
			"m": p,
		})
	})
	booksRouter.GET("/books", controllers.FindBooks)
	booksRouter.GET("/books/bulk-edit", controllers.BulkEdit)
	booksRouter.GET("/books/:id", controllers.FindBook)
	booksRouter.POST("/books", controllers.CreateBook)
	booksRouter.PATCH("/books/:id", controllers.UpdateBook)
	booksRouter.DELETE("/books/:id", controllers.DeleteBook)
}
