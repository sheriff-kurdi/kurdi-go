package repositories

import (
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/responses"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	repository := BookRepository{}
	return &repository
}
func (repository BookRepository) ListAllDiscounted() (discountedBooks []responses.BookResponse, err error) {
	err = database.PostgresDB.Model(&models.Book{}).Where("is_discounted = ?", true).Find(&discountedBooks).Error
	return
}

func (repository BookRepository) ListAllDiscountedModel() (discountedBooks []models.Book, err error) {
	err = database.PostgresDB.Model(&models.Book{}).Where("is_discounted = ?", true).Find(&discountedBooks).Error
	return
}
