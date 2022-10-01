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

	booksApp.Get("/books", controllers.Find)
	booksApp.Get("/books/bulk-edit", controllers.BulkEdit)
	booksApp.Get("/books/:id", controllers.FindById)
	booksApp.Post("/books", controllers.Create)
	booksApp.Patch("/books/:id", controllers.Update)
	booksApp.Delete("/books/:id", controllers.Delete)
}
