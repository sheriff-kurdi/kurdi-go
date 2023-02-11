package services

import (
	requests "kurdi-go/src/api/requests/products"
	resources "kurdi-go/src/api/responses/base"
	productsAggregate "kurdi-go/src/domain/entities/aggregates/products"
	"kurdi-go/src/infrastructure/database"
)

type ProductsService struct {
}

func NewProductsService() *ProductsService {
	service := ProductsService{}
	return &service
}

func (service ProductsService) ListAll() (response resources.IResponse) {
	var products []productsAggregate.Product
	err := database.PostgresDB.Model(&productsAggregate.Product{}).Find(&products).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(products, "")
}

func (service ProductsService) FindBySKU(sku string) (response resources.IResponse) {
	var product productsAggregate.Product
	if err := database.PostgresDB.Model(productsAggregate.Product{}).Where("sku = ?", sku).First(&product).Error; err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(product, "")
}

func (service ProductsService) Create(request requests.CreateProductRequest) (response resources.IResponse) {
	product := productsAggregate.Product{}
	err := database.PostgresDB.Create(&product).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(product, "created")
}

func (service ProductsService) Update(request requests.UpdateProductRequest, SKU string) (response resources.IResponse) {
	product := productsAggregate.Product{}
	err := database.PostgresDB.Save(&product).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(product, "updated")
}

func (service ProductsService) Delete(SKU string) (response resources.IResponse) {
	err := database.PostgresDB.Delete(&productsAggregate.Product{}, SKU).Error
	if err != nil {
		return resources.GetError500Response(err.Error())
	}
	return resources.GetSuccess200Response(SKU, "Deleted")

}
