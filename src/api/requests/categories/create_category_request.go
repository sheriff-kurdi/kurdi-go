package categories

import "kurdi-go/src/domain/entities/aggregates/products"

type CreateCategoryRequest struct {
	SKU string `json:"sku"`
	products.ProductDetails
	products.ProductPrice
	products.ProductQuantity
}
