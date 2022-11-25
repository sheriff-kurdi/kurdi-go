package routes

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/api/controllers"
	"kurdi-go/api/middlewares"
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
