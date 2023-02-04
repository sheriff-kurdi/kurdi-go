package services

import (
	"kurdi-go/api/requests"
	resources "kurdi-go/api/responses/base"
	"kurdi-go/api/responses"
	"kurdi-go/domain/entities"
	"kurdi-go/infrastructure/database"
)

type BookService struct {
}

func NewBookService() *BookService {
	service := BookService{}
	return &service
}

func (service BookService) ListAll() (response resources.IResponse) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}

	return resources.GetSuccess200Response(books, "")
}

func (service BookService) FindById(bookId int) (response resources.IResponse) {
	var book responses.BookResponse
	if err := database.PostgresDB.Model(entities.Book{}).Where("id = ?", bookId).First(&book).Scan(&book).Error; err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(book, "")
}

func (service BookService) Create(request requests.BookRequest) (response resources.IResponse) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	bookModel := entities.Book{Author: request.Author, Title: request.Title}
	err = database.PostgresDB.Create(&bookModel).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(bookModel, "created")
}

func (service BookService) Update(request requests.BookRequest, bookId int) (response resources.IResponse) {
	var bookModel entities.Book
	bookModel.Title = request.Title
	bookModel.Author = request.Author
	bookModel.ID = uint(bookId)
	err := database.PostgresDB.Save(&bookModel).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(bookModel, "updated")
}

func (service BookService) Delete(bookId int) (response resources.IResponse) {
	err := database.PostgresDB.Delete(&entities.Book{}, bookId).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(bookId, "Deleted")

}
