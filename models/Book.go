package models

import (
	"gorm.io/gorm"
	"kurdi-go/responses"
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

func (book Book) GetBookResponse() responses.BookResponse {
	var bookResponse responses.BookResponse

	bookResponse.Id = int(book.ID)
	bookResponse.CreatedAt = book.CreatedAt
	bookResponse.UpdatedAt = book.UpdatedAt
	bookResponse.Title = book.Title
	bookResponse.Author = book.Author
	bookResponse.CostPrice = book.CostPrice
	bookResponse.SellingPrice = book.SellingPrice
	bookResponse.Discount = book.Discount
	bookResponse.IsDiscounted = book.IsDiscounted

	if book.IsDiscounted {
		bookResponse.SellingPrice = book.SellingPrice - book.Discount
	}
	return bookResponse
}
