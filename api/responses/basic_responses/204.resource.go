package basic_responses

import "github.com/gofiber/fiber/v2"

type NoContent204 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetNoContent204Resource(data interface{}, message string) IResource {
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	}
	resource := NoContent204{Status: 204, Message: "", Data: dataJson}
	return &resource
}

func (NoContent204 NoContent204) GetStatus() int {
	return NoContent204.Status
}

func (NoContent204 NoContent204) GetData() *fiber.Map {
	return NoContent204.Data
}
