package controllers

import (
	"github.com/gin-gonic/gin"
	"kurdi-go/database"
	"kurdi-go/models"
	"kurdi-go/requests"
	"net/http"
)

// FindBooks GET /books
// Find all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	database.PostgresDB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// BulkEdit GET /books/bulk-edit
// Bulk Edit all books
func BulkEdit(c *gin.Context) {

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
		c.JSON(http.StatusOK, gin.H{"data": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": rows})
}

// FindBook GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook POST /books
// Create new book
func CreateBook(c *gin.Context) {
	var request requests.BookRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: request.Title, Author: request.Author}
	database.PostgresDB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate request
	var request requests.BookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = request.Title
	book.Author = request.Author
	database.PostgresDB.Save(request)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := database.PostgresDB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.PostgresDB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
