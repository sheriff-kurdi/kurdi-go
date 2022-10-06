package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/requests"
	"kurdi-go/resources"
	"kurdi-go/services"
	"strconv"
)

type BookController struct {
	service *services.BookService
}

func NewBookController() *BookController {
	controller := BookController{
		service: services.NewBookService(),
	}
	return &controller
}

// GetAll GET /books
// GetAll all books
func (controller BookController) GetAll(ctx *fiber.Ctx) error {
	//get all
	response := controller.service.ListAll()
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// GetAllDiscounted GET /books/discounted
// GetAllDiscounted all discounted books
func (controller BookController) GetAllDiscounted(ctx *fiber.Ctx) error {
	//get all
	response := controller.service.ListAllDiscounted()
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// FindById GET /books/:id
// Find a book
func (controller BookController) FindById(ctx *fiber.Ctx) error {
	bookId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//find by id
	response := controller.service.FindById(bookId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Create POST /books
// Create new book
func (controller BookController) Create(ctx *fiber.Ctx) error {
	var request requests.BookRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//create
	response := controller.service.Create(request)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Update PATCH /books/:id
// Update a book
func (controller BookController) Update(ctx *fiber.Ctx) error {
	//check param
	bookId, err := strconv.Atoi(ctx.Params("id"))
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
	response := controller.service.Update(request, bookId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Delete DELETE /books/:id
// Delete a book
func (controller BookController) Delete(ctx *fiber.Ctx) error {
	//check param
	bookId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//delete
	response := controller.service.Delete(bookId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
