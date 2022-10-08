package responses

import (
	"kurdi-go/models"
	"time"
)

type BookResponse struct {
	Id           int       `json:"id" binding:"id"`
	Title        string    `json:"title" binding:"required"`
	Author       string    `json:"author" binding:"required"`
	CostPrice    float32   `json:"cost_price"`
	SellingPrice float32   `json:"selling_price"`
	Discount     float32   `json:"discount"`
	IsDiscounted bool      `json:"is_discounted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (bookResponse BookResponse) ToResponse(bookModel models.Book) BookResponse {
	bookResponse.Id = int(bookModel.ID)
	bookResponse.Title = bookModel.Title
	bookResponse.Author = bookModel.Author
	bookResponse.CostPrice = bookModel.CostPrice
	bookResponse.SellingPrice = bookModel.SellingPrice
	bookResponse.Discount = bookModel.Discount
	bookResponse.IsDiscounted = bookModel.IsDiscounted
	bookResponse.CreatedAt = bookModel.CreatedAt
	bookResponse.UpdatedAt = bookModel.UpdatedAt
	return bookResponse
}
