package basic_responses

import "github.com/gofiber/fiber/v2"

type Error403 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError403Resource(message string) IResource {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	resource := Error403{Status: 403, Message: message, Data: dataJson}
	return &resource
}

func (error403 Error403) GetStatus() int {
	return error403.Status
}

func (error403 Error403) GetData() *fiber.Map {
	return error403.Data
}
