package categories

import (
	"kurdi-go/src/domain/entities"
	"kurdi-go/src/domain/entities/aggregates/categories"
)

type CategoryResponse struct {
	Name     string           `json:"name"`
	Details categories.CategoryDetails `json:"details"`
	entities.Entity
}
