package responses

import "kurdi-go/models"

type BooksResponse struct {
	books []BookResponse
}

func (booksResponse *BooksResponse) ToResponse(booksModel []models.Book) []BookResponse {
	for _, bookModel := range booksModel {
		bookResponse := BookResponse{}
		booksResponse.books = append(booksResponse.books, bookResponse.ToResponse(bookModel))
	}
	return booksResponse.books
}
