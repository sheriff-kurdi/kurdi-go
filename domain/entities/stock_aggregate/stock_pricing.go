package entities_stock_aggregate

type StockPricing struct {
	CostPrice    float32 `json:"cost_price"`
	SellingPrice float32 `json:"selling_price"`
	IsDiscounted bool    `json:"is_discounted"`
	Discount     float32 `json:"discount"`
}
