package base

import "github.com/gofiber/fiber/v2"

type Error403 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError403Response(message string) IResponse {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	response := Error403{Status: 403, Message: message, Data: dataJson}
	return &response
}

func (error403 Error403) GetStatus() int {
	return error403.Status
}

func (error403 Error403) GetData() *fiber.Map {
	return error403.Data
}
