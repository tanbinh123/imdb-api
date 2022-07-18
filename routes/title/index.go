package title

import (
	"errors"

	"github.com/Scrip7/imdb-api/services/title/index"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	// TODO: better group validation
	if err != nil {
		return errors.New("title id should be numerical, without 'tt' at the beginning")
	}

	res, err := index.Title(id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
