package basic_responses

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"kurdi-go/infrastructure"
)

type Error500 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError500Resource(message string) IResource {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	infrastructure.Info(json.Marshal(errors))
	resource := Error500{Status: 500, Message: "", Data: dataJson}
	return &resource
}

func (error500 Error500) GetStatus() int {
	return error500.Status
}

func (error500 Error500) GetData() *fiber.Map {
	return error500.Data
}
