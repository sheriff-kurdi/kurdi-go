package database

import (
	"kurdi-go/src/domain/entities"
	models "kurdi-go/src/domain/entities/aggregates/products"
)

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&entities.Book{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&models.Product{})
	if err != nil {
		return
	}
}

func SQLLiteAutoMigrate() {
	err := SqliteDB.AutoMigrate(&entities.Book{})
	if err != nil {
		return
	}
}
