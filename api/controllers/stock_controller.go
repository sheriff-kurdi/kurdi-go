package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests_stock "kurdi-go/api/requests/stock"
	"kurdi-go/api/responses/basic_responses"
	"kurdi-go/services"
	lib "kurdi-go/shared"
	"strconv"
)

type StockController struct {
	service *services.StockService
}

func NewStockController() *StockController {
	controller := StockController{
		service: services.NeStockService(),
	}
	return &controller
}

// GetAll GET /stock
// GetAll all stock
func (controller StockController) GetAll(ctx *fiber.Ctx) error {
	//get all
	response := controller.service.ListAll()
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// FindById GET /stocks/:id
// Find a stock
func (controller StockController) FindById(ctx *fiber.Ctx) error {
	sku := ctx.Params("id")
	//find by id
	response := controller.service.FindById(sku)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Create POST /stock
// Create new stock
func (controller StockController) Create(ctx *fiber.Ctx) error {
	var request requests_stock.CreateStockRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//create
	response := controller.service.Create(request)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Update PUT /stock/:id
// Update a stock
func (controller StockController) Update(ctx *fiber.Ctx) error {
	//check param
	sku := ctx.Params("sku")
	// Validate request
	var request requests_stock.UpdateStockRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	if sku != request.SKU {
		response := basic_responses.GetError400Resource(lib.Translate("GENERAL.THERE_IS_AN_ERROR"))
		return ctx.Status(response.GetStatus()).JSON(response.GetData())

	}
	//update
	response := controller.service.Update(request, sku)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Reserve PUT /stock/reserve/:id
// Update a stock
func (controller StockController) Reserve(ctx *fiber.Ctx) error {
	//check param
	sku := ctx.Params("sku")
	// Validate request
	var request requests_stock.QuantityRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//update
	response := controller.service.Reserve(request, sku)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// UnReserve PUT /stock/un-reserve/:id
// Update a stock
func (controller StockController) UnReserve(ctx *fiber.Ctx) error {
	//check param
	sku := ctx.Params("sku")
	// Validate request
	var request requests_stock.QuantityRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//update
	response := controller.service.UnReserve(request, sku)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// AddStock PUT /stock/add-stock/:id
// Update a stock
func (controller StockController) AddStock(ctx *fiber.Ctx) error {
	//check param
	sku := ctx.Params("sku")
	// Validate request
	var request requests_stock.QuantityRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//update
	response := controller.service.AddStock(request, sku)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Delete DELETE /stock/:id
// Delete a book
func (controller StockController) Delete(ctx *fiber.Ctx) error {
	//check param
	stockId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := basic_responses.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//delete
	response := controller.service.Delete(stockId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
