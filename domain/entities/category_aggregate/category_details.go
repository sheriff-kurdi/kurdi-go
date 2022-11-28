package entities_category_aggregate

import "kurdi-go/domain/entities"

type CategoryDetails struct {
	Id           string            `json:"id" gorm:"primaryKey"`
	Name         string            `json:"name" gorm:"primaryKey"`
	LanguageCode string            `json:"language_code" gorm:"primaryKey"`
	Language     entities.Language `gorm:"foreignKey:language_code;references:code"`
	entities.Entity
}
