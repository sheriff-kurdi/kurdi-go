package resources

import "github.com/gofiber/fiber/v2"

type IResource interface {
	GetStatus() int
	GetData() *fiber.Map
}
