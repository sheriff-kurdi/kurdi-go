package services

import (
	"kurdi-go/api/requests"
	resources2 "kurdi-go/api/resources"
	"kurdi-go/api/responses"
	"kurdi-go/domain/entities"
	entities_stock_aggregate "kurdi-go/domain/entities/stock_aggregate"
	"kurdi-go/infrastructure/database"
)

type StockService struct {
}

func NeStockService() *StockService {
	service := StockService{}
	return &service
}

func (service StockService) ListAll() (response resources2.IResource) {
	var stockItems []entities_stock_aggregate.StockItem
	err := database.PostgresDB.Model(&entities_stock_aggregate.StockItem{}).Scan(&stockItems).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockItems, "")
}

func (service StockService) FindById(stockId int) (response resources2.IResource) {
	var stockItem entities_stock_aggregate.StockItem
	if err := database.PostgresDB.Model(entities_stock_aggregate.StockItem{}).Where("id = ?", stockId).First(&stockItem).Scan(&stockItem).Error; err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockItem, "")
}

func (service StockService) Create(request requests.BookRequest) (response resources2.IResource) {
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

func (service StockService) Update(request requests.BookRequest, bookId int) (response resources2.IResource) {
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

func (service StockService) Delete(stockId int) (response resources2.IResource) {
	err := database.PostgresDB.Delete(&entities_stock_aggregate.StockItem{}, stockId).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockId, "Deleted")

}
