package services

import (
	"kurdi-go/api/requests"
	"kurdi-go/api/responses"
	"kurdi-go/api/responses/basic_responses"
	"kurdi-go/domain/entities"
	"kurdi-go/infrastructure/database"
)

type BookService struct {
}

func NewBookService() *BookService {
	service := BookService{}
	return &service
}

func (service BookService) ListAll() (response basic_responses.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}

	return basic_responses.GetSuccess200Resource(books, "")
}

func (service BookService) FindById(bookId int) (response basic_responses.IResource) {
	var book responses.BookResponse
	if err := database.PostgresDB.Model(entities.Book{}).Where("id = ?", bookId).First(&book).Scan(&book).Error; err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(book, "")
}

func (service BookService) Create(request requests.BookRequest) (response basic_responses.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	bookModel := entities.Book{Author: request.Author, Title: request.Title}
	err = database.PostgresDB.Create(&bookModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(bookModel, "created")
}

func (service BookService) Update(request requests.BookRequest, bookId int) (response basic_responses.IResource) {
	var bookModel entities.Book
	bookModel.Title = request.Title
	bookModel.Author = request.Author
	bookModel.ID = bookId
	err := database.PostgresDB.Save(&bookModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(bookModel, "updated")
}

func (service BookService) Delete(bookId int) (response basic_responses.IResource) {
	err := database.PostgresDB.Delete(&entities.Book{}, bookId).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(bookId, "Deleted")

}
