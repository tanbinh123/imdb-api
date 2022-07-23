package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// HTTPError is a API response wrapper
type HTTPError struct {
	OK      bool   `json:"ok"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Err(err).Str("path", c.Path()).Msg("Fiber error handler")

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(HTTPError{
		OK:      false,
		Code:    code,
		Message: err.Error(),
	})
}
