package services

import (
	"kurdi-go/src/api/requests"
	"kurdi-go/src/api/responses"
	resources "kurdi-go/src/api/responses/base"
	"kurdi-go/src/domain/entities"
	models "kurdi-go/src/domain/entities/aggregates/products"
	"kurdi-go/src/infrastructure/database"
)

type StockService struct {
}

func NeStockService() *StockService {
	service := StockService{}
	return &service
}

func (service StockService) ListAll() (response resources.IResponse) {
	var stockItems []models.Product
	err := database.PostgresDB.Model(&models.Product{}).Where("sa").First(&stockItems).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(stockItems, "")
}

func (service StockService) FindById(bookId int) (response resources.IResponse) {
	var stock responses.BookResponse
	if err := database.PostgresDB.Model(entities.Book{}).Where("id = ?", bookId).First(&stock).Error; err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(stock, "")
}

func (service StockService) Create(request requests.BookRequest) (response resources.IResponse) {
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

func (service StockService) Update(request requests.BookRequest, bookId int) (response resources.IResponse) {
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

func (service StockService) Delete(bookId int) (response resources.IResponse) {
	err := database.PostgresDB.Delete(&entities.Book{}, bookId).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(bookId, "Deleted")

}
