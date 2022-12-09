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
	stockApp.Put("/stock/:id", stockController.Update)
	stockApp.Put("/stock/reserve/:id", stockController.Reserve)
	stockApp.Put("/stock/un-reserve/:id", stockController.UnReserve)
	stockApp.Put("/stock/add-stock/:id", stockController.AddStock)
	stockApp.Delete("/stock/:id", stockController.Delete)
}
