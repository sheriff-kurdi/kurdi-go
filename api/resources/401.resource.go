package resources

import "github.com/gofiber/fiber/v2"

type Error401 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetError401Resource(message string) IResource {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	resource := Error401{Status: 401, Message: message, Data: dataJson}
	return &resource
}

func (error401 Error401) GetStatus() int {
	return error401.Status
}

func (error401 Error401) GetData() *fiber.Map {
	return error401.Data
}
