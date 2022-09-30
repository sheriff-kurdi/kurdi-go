package routes

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/controllers"
	"kurdi-go/middlewares"
)

func BooksRoutes(booksApp *fiber.App) {
	booksApp.Use(middlewares.AuthenticationMiddleware)
	booksApp.Get("/get", func(c *fiber.Ctx) error {
		var result = 5
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    result,
		})
	})

	booksApp.Get("/books", controllers.FindBooks)
	booksApp.Get("/books/bulk-edit", controllers.BulkEdit)
	booksApp.Get("/books/:id", controllers.FindBook)
	booksApp.Post("/books", controllers.CreateBook)
	booksApp.Patch("/books/:id", controllers.UpdateBook)
	booksApp.Delete("/books/:id", controllers.DeleteBook)
}
