package routes

import (
	"kurdi-go/src/api/controllers"
	"kurdi-go/src/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(stockApp *fiber.App) {
	stockController := controllers.NewStockController()
	stockApp.Use(middlewares.AuthenticationMiddleware)

	stockApp.Get("/stock", stockController.GetAll)
	stockApp.Get("/stock/:id", stockController.FindById)
	stockApp.Post("/stock", stockController.Create)
	stockApp.Patch("/stock/:id", stockController.Update)
	stockApp.Delete("/stock/:id", stockController.Delete)
}
