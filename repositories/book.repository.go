package repositories

import (
	"gorm.io/gorm"
	"kurdi-go/models"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	repository := BookRepository{}
	return &repository
}

func (repository BookRepository) ListAllDiscounted(connection *gorm.DB) (discountedBooks []models.Book, err error) {
	err = connection.Model(&models.Book{}).Where("is_discounted = ?", true).Find(&discountedBooks).Error
	return
}

func (repository BookRepository) ListAll(connection *gorm.DB) (Books []models.Book, err error) {
	err = connection.Model(&models.Book{}).Find(&Books).Error
	return
}
