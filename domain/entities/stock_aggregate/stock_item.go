package entities_stock_aggregate

import "kurdi-go/domain/entities"

type StockItem struct {
	SKU     string         `json:"sku" gorm:"primaryKey"`
	Details []StockDetails `json:"Details" gorm:"foreignKey:sku;references:sku"`
	StockPricing
	StockQuantity
	entities.Entity
}



func (stock *StockItem) Reserve(reservationQuantity int) (*StockItem) {

	stock.AvailableStock -= reservationQuantity
	stock.ReservedStock += reservationQuantity
	return stock
}

func (stock *StockItem) UnReserve(reservationQuantity int) (*StockItem) {

	stock.AvailableStock += reservationQuantity
	stock.ReservedStock -= reservationQuantity
	return stock
}

func (stock *StockItem) AddStock(adedStock int) (*StockItem) {

	stock.AvailableStock += adedStock
	return stock
}