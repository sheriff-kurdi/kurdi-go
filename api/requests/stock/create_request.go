package requests_stock

import entities_stock_aggregate "kurdi-go/domain/entities/stock_aggregate"

type CreateStockRequest struct {
	SKU          string                                  `json:"sku" gorm:"primaryKey"`
	Details      []entities_stock_aggregate.StockDetails `json:"Details" gorm:"foreignKey:sku;references:sku"`
	TotalStock   int                                     `json:"total_stock"`
	SellingPrice float32                                 `json:"selling_price"`
	IsDiscounted bool                                    `json:"is_discounted"`
	Discount     float32                                 `json:"discount"`
}

type StockDetailsRequest struct {
	SKU          string `json:"sku" gorm:"primaryKey"`
	LanguageCode string `json:"language_code" gorm:"primaryKey"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
