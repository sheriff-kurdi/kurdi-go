package test_books

import (
	"github.com/stretchr/testify/assert"
	"kurdi-go/math"
	"testing"
)

func ListAllDiscountedTest(test *testing.T) {
	// preparation
	/*	booksService := services.BookService{}
		books := booksService.ListAllDiscounted()*/

	assert := assert.New(test)

	actual := math.Add(1, 2)
	expected := 3

	assert.Equal(expected, actual, "The two numbers should be equals.")
}
