package products

type ProductQuantity struct {
	TotalStock     int `json:"total_stock"`
	AvailableStock int `json:"available_stock"`
	ReservedStock  int `json:"reserved_stock"`
}
