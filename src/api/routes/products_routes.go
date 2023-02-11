package routes

import (
	"kurdi-go/src/api/controllers"
	"kurdi-go/src/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(productsRoutes *fiber.App) {
	productsController := controllers.NewProductsController()
	productsRoutes.Use(middlewares.AuthenticationMiddleware)

	productsRoutes.Get("/stock", productsController.GetAll)
	productsRoutes.Get("/stock/:id", productsController.FindById)
	productsRoutes.Post("/stock", productsController.Create)
	productsRoutes.Patch("/stock/:id", productsController.Update)
	productsRoutes.Delete("/stock/:id", productsController.Delete)
}
