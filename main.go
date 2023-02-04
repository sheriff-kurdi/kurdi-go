package main

import (
	"kurdi-go/src/api/routes"
	"kurdi-go/src/infrastructure/database"
	"kurdi-go/src/infrastructure/logger"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//app
	app := fiber.New()
	logger := logger.ZapLogger()

	if err := godotenv.Load(".env"); err != nil {
		logger.Warn("Error loading .env testing file")
	}

	// Connect to Database
	database.Connect()
	//Migrate Database
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	routes.StockRoutes(app)

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("App started!")
}
