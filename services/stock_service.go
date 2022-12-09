package services

import (
	requests_stock "kurdi-go/api/requests/stock"
	"kurdi-go/api/responses/basic_responses"
	entities_stock_aggregate "kurdi-go/domain/entities/stock_aggregate"
	"kurdi-go/infrastructure/database"
)

type StockService struct {
}

func NeStockService() *StockService {
	service := StockService{}
	return &service
}

func (service StockService) ListAll() (response basic_responses.IResource) {
	var stockItems []entities_stock_aggregate.StockItem
	err := database.PostgresDB.Model(&entities_stock_aggregate.StockItem{}).Preload("Details", "language_code = ?", "ar").Find(&stockItems).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItems, "")
}

func (service StockService) FindById(stockId string) (response basic_responses.IResource) {
	var stockItem entities_stock_aggregate.StockItem
	if err := database.PostgresDB.Model(entities_stock_aggregate.StockItem{}).Where("sku = ?", stockId).First(&stockItem).Preload("Details", "language_code = ?", "ar").Find(&stockItem).Error; err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItem, "")
}

func (service StockService) Create(request requests_stock.CreateStockRequest) (response basic_responses.IResource) {
	var stockItemModel entities_stock_aggregate.StockItem
	stockItemModel.SKU = request.SKU
	stockItemModel.Details = request.Details
	stockItemModel.TotalStock = request.TotalStock
	stockItemModel.SellingPrice = request.SellingPrice
	stockItemModel.Discount = request.Discount
	stockItemModel.IsDiscounted = request.IsDiscounted

	err := database.PostgresDB.Save(&stockItemModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItemModel, "created")
}

func (service StockService) Update(request requests_stock.UpdateStockRequest, sku string) (response basic_responses.IResource) {
	var stockItemModel entities_stock_aggregate.StockItem
	stockItemModel.SKU = request.SKU
	stockItemModel.Details = request.Details
	stockItemModel.TotalStock = request.TotalStock
	stockItemModel.SellingPrice = request.SellingPrice
	stockItemModel.Discount = request.Discount
	stockItemModel.IsDiscounted = request.IsDiscounted

	err := database.PostgresDB.Save(&stockItemModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItemModel, "updated")
}

func (service StockService) Reserve(request requests_stock.QuantityRequest, sku string) (response basic_responses.IResource) {

	var stockItemModel entities_stock_aggregate.StockItem
	if database.PostgresDB.Where("sku = ?", sku).Find(&stockItemModel).Error != nil{
		return basic_responses.GetError500Resource("not found")
	}

	stockItemModel.Reserve(request.Quantity)

	err := database.PostgresDB.Save(&stockItemModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItemModel, "updated")
}

func (service StockService) UnReserve(request requests_stock.QuantityRequest, sku string) (response basic_responses.IResource) {

	var stockItemModel entities_stock_aggregate.StockItem
	if database.PostgresDB.Where("sku = ?", sku).Find(&stockItemModel).Error != nil{
		return basic_responses.GetError500Resource("not found")
	}

	stockItemModel.UnReserve(request.Quantity)

	err := database.PostgresDB.Save(&stockItemModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItemModel, "updated")
}

func (service StockService) AddStock(request requests_stock.QuantityRequest, sku string) (response basic_responses.IResource) {

	var stockItemModel entities_stock_aggregate.StockItem
	if database.PostgresDB.Where("sku = ?", sku).Find(&stockItemModel).Error != nil{
		return basic_responses.GetError500Resource("not found")
	}

	stockItemModel.AddStock(request.Quantity)

	err := database.PostgresDB.Save(&stockItemModel).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockItemModel, "updated")
}

func (service StockService) Delete(stockId int) (response basic_responses.IResource) {
	err := database.PostgresDB.Delete(&entities_stock_aggregate.StockItem{}, stockId).Error
	if err != nil {
		return basic_responses.GetError500Resource(err.Error())
	}
	return basic_responses.GetSuccess200Resource(stockId, "Deleted")

}
