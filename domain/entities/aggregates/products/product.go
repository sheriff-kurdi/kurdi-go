package products

import "kurdi-go/domain/entities"

type Product struct {
	SKU string `json:"sku"`
	ProductDetails
	ProductPrice
	ProductQuantity
	entities.Entity
}
