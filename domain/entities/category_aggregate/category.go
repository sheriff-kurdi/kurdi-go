package entities_category_aggregate

import "kurdi-go/domain/entities"

type Category struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	IsParent       bool      `json:"is_parent"`
	ParentCategory string    `json:"parent_category"`
	Parent         *Category `json:"parent" gorm:"foreignKey:parent_category;references:id"`
	entities.Entity
}
