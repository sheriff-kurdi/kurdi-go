package base

import "github.com/gofiber/fiber/v2"

type Success200 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetSuccess200Response(data interface{}, message string) IResponse {

	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	}

	response := Success200{Status: 200, Message: "", Data: dataJson}
	return &response
}

func (success200 Success200) GetStatus() int {
	return success200.Status
}

func (success200 Success200) GetData() *fiber.Map {
	return success200.Data
}
