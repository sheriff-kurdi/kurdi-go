package products

import "kurdi-go/src/domain/entities/aggregates/products"

type CreateProductRequest struct {
	SKU string `json:"sku"`
	products.ProductDetails
	products.ProductPrice
	products.ProductQuantity
}
