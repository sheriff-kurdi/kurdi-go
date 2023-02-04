package base

import "github.com/gofiber/fiber/v2"

type IResponse interface {
	GetStatus() int
	GetData() *fiber.Map
}
