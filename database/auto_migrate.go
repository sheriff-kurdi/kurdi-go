package database

import "kurdi-go/models"

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&models.Book{})
	if err != nil {
		return
	}
}
