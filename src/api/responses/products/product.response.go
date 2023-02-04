package products

import (
	models "kurdi-go/src/domain/entities/aggregates/products"
)

type ProductResponse struct {
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	models.ProductPrice
	models.ProductQuantity
}
