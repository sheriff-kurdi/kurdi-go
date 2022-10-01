package routes

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/controllers"
	"kurdi-go/middlewares"
)

func BooksRoutes(booksApp *fiber.App) {
	bookController := controllers.NewBookController()
	booksApp.Use(middlewares.AuthenticationMiddleware)
	booksApp.Get("/get", func(c *fiber.Ctx) error {
		var result = 5
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    result,
		})
	})

	booksApp.Get("/books", bookController.GetAll)
	booksApp.Get("/books/:id", bookController.FindById)
	booksApp.Post("/books", bookController.Create)
	booksApp.Patch("/books/:id", bookController.Update)
	booksApp.Delete("/books/:id", bookController.Delete)
}
