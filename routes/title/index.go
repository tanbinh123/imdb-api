package title

import (
	"github.com/Scrip7/imdb-api/services/title/index"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	res, err := index.Title(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
