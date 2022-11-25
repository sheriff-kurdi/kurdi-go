package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/api/requests"
	"kurdi-go/api/resources"
	"kurdi-go/services"
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
	stockId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//find by id
	response := controller.service.FindById(stockId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Create POST /stock
// Create new stock
func (controller StockController) Create(ctx *fiber.Ctx) error {
	var request requests.BookRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := resources.GetError500Resource(err.Error())
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
	stockId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	// Validate request
	var request requests.BookRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//update
	response := controller.service.Update(request, stockId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Delete DELETE /stock/:id
// Delete a book
func (controller StockController) Delete(ctx *fiber.Ctx) error {
	//check param
	stockId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//delete
	response := controller.service.Delete(stockId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
