package title

import (
	"errors"

	"github.com/Scrip7/imdb-api/services"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	// TODO: better group validation
	if err != nil {
		return errors.New("title id should be numerical, without 'tt' at the beginning")
	}

	res, err := services.Title(id)
	if err != nil {
		// TODO: global error handling middleware
		return err
	}
	return c.JSON(res)
}
