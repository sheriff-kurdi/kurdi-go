package responses

import (
	entities_stock_aggregate "kurdi-go/domain/entities/stock_aggregate"
)

type StockResponse struct {
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	entities_stock_aggregate.StockPricing
	entities_stock_aggregate.StockQuantity
}
