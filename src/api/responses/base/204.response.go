package base

import "github.com/gofiber/fiber/v2"

type NoContent204 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetNoContent204Response(data interface{}, message string) IResponse {
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	}
	response := NoContent204{Status: 204, Message: "", Data: dataJson}
	return &response
}

func (NoContent204 NoContent204) GetStatus() int {
	return NoContent204.Status
}

func (NoContent204 NoContent204) GetData() *fiber.Map {
	return NoContent204.Data
}
