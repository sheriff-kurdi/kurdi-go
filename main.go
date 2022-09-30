package main

import (
	"github.com/gin-gonic/gin"
	"kurdi-go/database"
	"kurdi-go/routes"
)

func main() {

	//Database
	// Connect to database
	database.Connect()
	//migrate database
	//TODO:Migrate based on env variable
	database.PostgresAutoMigrate()
	database.SQLLiteAutoMigrate()

	// Routes
	router := gin.Default()
	routes.BooksRoutes(router)

	// Run the server
	err := router.Run()
	if err != nil {
		return
	}
}
