package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/requests"
)

// FindBooks GET /books
// Find all books
func FindBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.PostgresDB.Find(&books)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    "books",
	})
}

// BulkEdit GET /books/bulk-edit
// Bulk Edit all books
func BulkEdit(c *fiber.Ctx) error {

	var books []models.Book
	database.PostgresDB.Find(&books)

	for _, book := range books {
		book.Title = "oo"
	}
	books[0].Title = "Database"

	booksJson := "[\n\t\t{\n\t\t\t\"ID\": 0,\n\t\t\t\"CreatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"UpdatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"DeletedAt\": null,\n\t\t\t\"id\": 1,\n\t\t\t\"title\": \"1\",\n\t\t\t\"author\": \"rr\"\n\t\t},\n\t\t{\n\t\t\t\"ID\": 0,\n\t\t\t\"CreatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"UpdatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"DeletedAt\": null,\n\t\t\t\"id\": 2,\n\t\t\t\"title\": \"2\",\n\t\t\t\"author\": \"tt\"\n\t\t},\n\t\t{\n\t\t\t\"ID\": 0,\n\t\t\t\"CreatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"UpdatedAt\": \"0001-01-01T00:00:00Z\",\n\t\t\t\"DeletedAt\": null,\n\t\t\t\"id\": 3,\n\t\t\t\"title\": \"3\",\n\t\t\t\"author\": \"yy\"\n\t\t}\n\t]"

	bulkUpdateQuery := `update books 
							set title = temp.title
							FROM (SELECT * FROM json_to_recordset(?) AS X("title" varchar, "id" int)) AS temp
							WHERE books.id = temp.id`
	rows, err := database.PostgresDB.Raw(bulkUpdateQuery, booksJson).Rows()
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "books",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    rows,
	})
}

// FindBook GET /books/:id
// Find a book
func FindBook(c *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.QueryParser("id")).First(&book).Error; err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "rows",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    "rows",
	})
}

// CreateBook POST /books
// Create new book
func CreateBook(c *fiber.Ctx) error {
	var request requests.BookRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "rows",
		})
	}

	// Create book
	book := models.Book{Title: request.Title, Author: request.Author}
	database.PostgresDB.Create(&book)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    "rows",
	})
}

// UpdateBook PATCH /books/:id
// Update a book
func UpdateBook(c *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.QueryParser("id")).First(&book).Error; err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "rows",
		})
	}

	// Validate request
	var request requests.BookRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "rows",
		})
	}

	book.Title = request.Title
	book.Author = request.Author
	database.PostgresDB.Save(request)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    "rows",
	})
}

// DeleteBook DELETE /books/:id
// Delete a book
func DeleteBook(c *fiber.Ctx) error {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.QueryParser("id")).First(&book).Error; err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    "rows",
		})
	}

	database.PostgresDB.Delete(&book)

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    "rows",
	})
}
