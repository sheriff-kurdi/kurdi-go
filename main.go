package main

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/database"
	"kurdi-go/routes"
)

func main() {
	//app
	app := fiber.New()

	//Database
	// Connect to database
	database.Connect()
	//migrate database
	//TODO:Migrate based on env variable
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	routes.BooksRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
