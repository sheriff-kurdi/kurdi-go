package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	CostPrice    float32 `json:"cost_price"`
	SellingPrice float32 `json:"selling_price"`
	Discount     float32 `json:"discount"`
	IsDiscounted bool    `json:"is_discounted"`
}

func (book *Book) AfterFind(tx *gorm.DB) (err error) {
	if book.IsDiscounted {
		book.SellingPrice -= book.Discount
	}
	return
}
