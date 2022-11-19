package database

import (
	"kurdi-go/domain/entities"
)

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&entities.Book{})
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
