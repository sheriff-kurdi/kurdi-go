package requests

type BookRequest struct {
	Title        string  `json:"title" binding:"required"`
	Author       string  `json:"author" binding:"required"`
	CostPrice    float32 `json:"cost_price"`
	SellingPrice float32 `json:"selling_price"`
	Discount     float32 `json:"discount"`
	IsDiscounted bool    `json:"is_discounted"`
}
