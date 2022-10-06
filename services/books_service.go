package services

import (
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/requests"
	"kurdi-go/resources"
	"kurdi-go/responses"
)

type BookService struct {
}

func NewBookService() *BookService {
	service := BookService{}
	return &service
}

func (service BookService) ListAll() (response resources.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&models.Book{}).Scan(&books).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(books, "")
}

func (service BookService) ListAllDiscounted() (response resources.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&models.Book{}).Where("is_discounted = ?", true).Scan(&books).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(books, "")
}

func (service BookService) FindById(bookId int) (response resources.IResource) {
	var book responses.BookResponse
	if err := database.PostgresDB.Model(models.Book{}).Where("id = ?", bookId).First(&book).Scan(&book).Error; err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(book, "")
}

func (service BookService) Create(request requests.BookRequest) (response resources.IResource) {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&models.Book{}).Scan(&books).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	bookModel := models.Book{Author: request.Author, Title: request.Title}
	err = database.PostgresDB.Create(&bookModel).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(bookModel, "created")
}

func (service BookService) Update(request requests.BookRequest, bookId int) (response resources.IResource) {
	var bookModel models.Book
	bookModel.Title = request.Title
	bookModel.Author = request.Author
	bookModel.ID = uint(bookId)
	err := database.PostgresDB.Save(&bookModel).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(bookModel, "updated")
}

func (service BookService) Delete(bookId int) (response resources.IResource) {
	err := database.PostgresDB.Delete(&models.Book{}, bookId).Error
	if err != nil {
		return resources.GetError500Resource(err.Error())
	}
	return resources.GetSuccess200Resource(bookId, "Deleted")

}
