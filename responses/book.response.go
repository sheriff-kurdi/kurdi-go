package responses

import (
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
