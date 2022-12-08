package entities_stock_aggregate

import "kurdi-go/domain/entities"

type StockDetails struct {
	SKU          string            `json:"sku" gorm:"primaryKey"`
	LanguageCode string            `json:"language_code" gorm:"primaryKey"`
	Language     entities.Language `json:"-" gorm:"foreignKey:language_code;references:code"`
	//StockItem    StockItem         `gorm:"foreignKey:sku;references:sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	entities.Entity
}
