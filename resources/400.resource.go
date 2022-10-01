package resources

import "github.com/gofiber/fiber/v2"

type Error400 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError400Resource(message string) IResource {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	resource := Error400{Status: 400, Message: message, Data: dataJson}
	return &resource
}

func (error400 Error400) GetStatus() int {
	return error400.Status
}

func (error400 Error400) GetData() *fiber.Map {
	return error400.Data
}
