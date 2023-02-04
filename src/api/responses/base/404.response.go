package base

import "github.com/gofiber/fiber/v2"

type Error404 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError404Response(message string) IResponse {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	response := Error404{Status: 404, Message: message, Data: dataJson}
	return &response
}

func (error404 Error404) GetStatus() int {
	return error404.Status
}

func (error404 Error404) GetData() *fiber.Map {
	return error404.Data
}
