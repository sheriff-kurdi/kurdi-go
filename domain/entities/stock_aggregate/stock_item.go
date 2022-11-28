package entities_stock_aggregate

import "kurdi-go/domain/entities"

type StockItem struct {
	SKU     string         `json:"sku" gorm:"primaryKey"`
	Details []StockDetails `json:"Details" gorm:"foreignKey:sku;references:sku"`
	StockPricing
	StockQuantity
	entities.Entity
}
