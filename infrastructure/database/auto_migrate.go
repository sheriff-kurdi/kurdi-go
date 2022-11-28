package database

import (
	"kurdi-go/domain/entities"
	entities_category_aggregate "kurdi-go/domain/entities/category_aggregate"
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
	err = PostgresDB.AutoMigrate(&entities_stock_aggregate.StockDetails{})
	if err != nil {
		return
	}

	err = PostgresDB.AutoMigrate(&entities_category_aggregate.Category{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&entities_category_aggregate.CategoryDetails{})
	if err != nil {
		return
	}

	err = PostgresDB.AutoMigrate(&entities.Language{})
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
