package categories

import (
	"kurdi-go/src/domain/entities"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Name     string           `json:"name"`
	Details CategoryDetails `json:"details"`
	entities.Entity
}

func (m *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Category) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	return
}

func (m *Category) AfterFind(tx *gorm.DB) (err error) {
	return
}
