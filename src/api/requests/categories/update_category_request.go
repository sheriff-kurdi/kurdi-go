package categories

import "kurdi-go/src/domain/entities/aggregates/products"

type UpdateCategoryRequest struct {
	SKU string `json:"sku"`
	products.ProductDetails
	products.ProductPrice
	products.ProductQuantity
}
