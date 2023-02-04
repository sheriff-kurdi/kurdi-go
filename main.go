package main

import (
	"fmt"
	"kurdi-go/src/api/routes"
	"kurdi-go/src/infrastructure/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//app
	app := fiber.New()

	envFile := godotenv.Load(".env")
	if envFile != nil {
		fmt.Print("Error loading .env testing file")
	}
	
	//Database
	// Connect to infrastructure
	database.Connect()
	//migrate infrastructure
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	routes.StockRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err := app.Listen(os.Getenv("PORT"))
	if err != nil {
		return
	}

}
