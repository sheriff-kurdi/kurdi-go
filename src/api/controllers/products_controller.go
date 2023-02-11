package controllers

import (
	requests "kurdi-go/src/api/requests/products"
	responses "kurdi-go/src/api/responses/base"
	"kurdi-go/src/services"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct {
	service *services.ProductsService
}

func NewProductsController() *ProductsController {
	controller := ProductsController{
		service: services.NewProductsService(),
	}
	return &controller
}

// GetAll GET /stock
// GetAll all stock
func (controller ProductsController) GetAll(ctx *fiber.Ctx) error {
	//get all
	response := controller.service.ListAll()
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// FindById GET /stocks/:id
// Find a stock
func (controller ProductsController) FindById(ctx *fiber.Ctx) error {
	stockId := ctx.Params("id")
	//find by id
	response := controller.service.FindBySKU(stockId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Create POST /stock
// Create new stock
func (controller ProductsController) Create(ctx *fiber.Ctx) error {
	var request requests.CreateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := responses.GetError500Response(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//create
	response := controller.service.Create(request)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Update PUT /stock/:id
// Update a stock
func (controller ProductsController) Update(ctx *fiber.Ctx) error {
	//check param
	SKU := ctx.Params("id")
	// Validate request
	var request requests.UpdateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := responses.GetError500Response(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//update
	response := controller.service.Update(request, SKU)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Delete DELETE /stock/:id
// Delete a book
func (controller ProductsController) Delete(ctx *fiber.Ctx) error {
	//check param
	SKU := ctx.Params("id")
	//delete
	response := controller.service.Delete(SKU)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
