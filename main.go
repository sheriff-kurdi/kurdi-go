package main

import (
	"kurdi-go/src/api/routes"
	"kurdi-go/src/infrastructure/database"

	"github.com/gofiber/fiber/v2"
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
	routes.StockRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
