package base

import (
	"kurdi-go/src/infrastructure/logger"

	"github.com/gofiber/fiber/v2"
)

type Error500 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError500Response(message string) IResponse {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": false,
		"message": "There is internal serval error.",
		"data":    errors,
	}
	logger := logger.ZapLogger()
	logger.Error(message)
	resource := Error500{Status: 500, Message: "", Data: dataJson}
	return &resource
}

func (error500 Error500) GetStatus() int {
	return error500.Status
}

func (error500 Error500) GetData() *fiber.Map {
	return error500.Data
}
