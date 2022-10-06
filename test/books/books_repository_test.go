package books_test

import (
	assertion "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/repositories"
	"testing"
)

func TestListOnlyDiscounted(test *testing.T) {
	assert := assertion.New(test)
	database.Connect()
	// preparation
	booksRepository := repositories.NewBookRepository()
	books, err := booksRepository.ListAllDiscounted()
	if err != nil {
		test.Error(err)
	}
	for _, book := range books {
		actual := book.IsDiscounted
		assert.True(actual, "The book should be discounted.")
	}
}

func TestListWithWriteSellingPrice(test *testing.T) {
	assert := assertion.New(test)
	database.Connect()
	// preparation
	booksRepository := repositories.NewBookRepository()
	books, err := booksRepository.ListAllDiscountedModel()
	if err != nil {
		test.Error(err)
	}
	for _, book := range books {
		var actualBook models.Book
		// scan neglect gorm hooks
		database.PostgresDB.Session(&gorm.Session{SkipHooks: true}).Model(models.Book{}).Where("id = ?", book.ID).First(&actualBook)
		if book.IsDiscounted {
			actual := book.SellingPrice
			expected := actualBook.SellingPrice - actualBook.Discount
			assert.Equal(expected, actual, "discounted price selling price should be calculated.")
		}

	}
}
