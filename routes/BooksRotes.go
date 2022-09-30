package routes

import (
	"github.com/gin-gonic/gin"
	"kurdi-go/controllers"
)

func BooksRoutes(router *gin.Engine) {
	router.GET("/get", func(context *gin.Context) {
		var a = 5
		var p = &a
		context.JSON(200, gin.H{
			"m": p,
		})
	})
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/bulk-edit", controllers.BulkEdit)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
