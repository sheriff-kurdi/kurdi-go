package main

import (
	"github.com/gin-gonic/gin"
	"kurdi-go/database"
	"kurdi-go/middlewares"
	"kurdi-go/routes"
)

func main() {

	router := gin.Default()
	router.Use(middlewares.CustomMiddleware)
	// Connect to database
	database.Connect()

	//migrate database
	//TODO:Migrate based on env variable
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	routes.BooksRoutes(router)

	// Run the server
	err := router.Run()
	if err != nil {
		return
	}
}
