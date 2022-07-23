package server

import "github.com/gofiber/fiber/v2"

func fallback(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Not found")
}
