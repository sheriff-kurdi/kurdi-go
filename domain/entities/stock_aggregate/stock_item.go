package entities_stock_aggregate

import "kurdi-go/domain/entities"

type StockItem struct {
	SKU string `json:"sku"`
	StockDetails
	StockPricing
	StockQuantity
	entities.Entity
}
