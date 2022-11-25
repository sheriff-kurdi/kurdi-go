package main

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/api/routes"
	"kurdi-go/infrastructure/database"
)

func main() {
	//app
	app := fiber.New()

	//Database
	// Connect to infrastructure
	database.Connect()
	//migrate infrastructure
	//TODO:Migrate based on env variable
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	routes.BooksRoutes(app)
	routes.StockRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
