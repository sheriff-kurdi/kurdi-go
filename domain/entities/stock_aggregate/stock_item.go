package stock_aggregate

import "kurdi-go/domain/entities"

type StockItem struct {
	entities.Entity
	SKU string `json:"sku"`
	StockDetails
	StockPricing
	StockQuantity
}
