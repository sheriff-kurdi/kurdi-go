package services

import (
	"kurdi-go/api/requests"
	resources2 "kurdi-go/api/resources"
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

func (service BookService) ListAll() (response resources2.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}

	return resources2.GetSuccess200Resource(books, "")
}

func (service BookService) FindById(bookId int) (response resources2.IResource) {
	var book responses.BookResponse
	if err := database.PostgresDB.Model(entities.Book{}).Where("id = ?", bookId).First(&book).Scan(&book).Error; err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(book, "")
}

func (service BookService) Create(request requests.BookRequest) (response resources2.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&entities.Book{}).Scan(&books).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	bookModel := entities.Book{Author: request.Author, Title: request.Title}
	err = database.PostgresDB.Create(&bookModel).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(bookModel, "created")
}

func (service BookService) Update(request requests.BookRequest, bookId int) (response resources2.IResource) {
	var bookModel entities.Book
	bookModel.Title = request.Title
	bookModel.Author = request.Author
	bookModel.ID = bookId
	err := database.PostgresDB.Save(&bookModel).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(bookModel, "updated")
}

func (service BookService) Delete(bookId int) (response resources2.IResource) {
	err := database.PostgresDB.Delete(&entities.Book{}, bookId).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(bookId, "Deleted")

}
