package resources

import "github.com/gofiber/fiber/v2"

type Success200 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetSuccess200Resource(data interface{}, message string) IResource {

	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	}

	resource := Success200{Status: 200, Message: "", Data: dataJson}
	return &resource
}

func (success200 Success200) GetStatus() int {
	return success200.Status
}

func (success200 Success200) GetData() *fiber.Map {
	return success200.Data
}
