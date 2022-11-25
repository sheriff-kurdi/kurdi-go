package database

import (
	"kurdi-go/domain/entities"
	entities_stock_aggregate "kurdi-go/domain/entities/stock_aggregate"
)

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&entities.Book{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&entities_stock_aggregate.StockItem{})
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
