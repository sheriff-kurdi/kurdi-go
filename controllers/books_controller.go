package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/requests"
	"kurdi-go/responses"
	"net/http"
)

type BookController struct {
}

func NewBookController() *BookController {
	controller := BookController{}
	return &controller
}

// GetAll GET /books
// GetAll all books
func (controller BookController) GetAll(ctx *fiber.Ctx) error {
	var books []responses.BookResponse
	err := database.PostgresDB.Model(&models.Book{}).Scan(&books).Error
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    books,
	})
}

// FindById GET /books/:id
// Find a book
func (controller BookController) FindById(ctx *fiber.Ctx) error {
	// Get model if exist
	var book responses.BookResponse
	id := ctx.Params("id")
	if err := database.PostgresDB.Model(models.Book{}).Where("id = ?", id).First(&book).Scan(&book).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    book,
	})
}

// Create POST /books
// Create new book
func (controller BookController) Create(ctx *fiber.Ctx) error {
	var request requests.BookRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	bookModel := models.Book{Author: request.Author, Title: request.Title}
	err := database.PostgresDB.Create(&bookModel).Error
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    bookModel,
	})
}

// Update PATCH /books/:id
// Update a book
func (controller BookController) Update(ctx *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	id := ctx.Params("id")
	if err := database.PostgresDB.Where("id = ?", id).First(&book).Error; err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	// Validate request
	var request requests.BookRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	book.Title = request.Title
	book.Author = request.Author
	database.PostgresDB.Save(&book)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    request,
	})
}

// Delete DELETE /books/:id
// Delete a book
func (controller BookController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", id).First(&book).Error; err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"message": "",
			"data":    err,
		})
	}

	database.PostgresDB.Delete(&book)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    book,
	})
}
