package handler

import (
	"github.com/Scrip7/imdb-api/pkg/title/crazycredits"
	"github.com/gofiber/fiber/v2"
)

// TitleCrazyCredits is a function that returns list of titles crazy credits
// @Summary Get titles crazy credits
// @Tags    Title
// @Param   id  path     string true "Title ID"
// @Success 200 {object} pipe.TitleCrazyCreditsTransform{}
// @Failure 404 {object} server.httpError
// @Failure 500 {object} server.httpError
// @Router  /title/{id}/crazycredits [get]
func TitleCrazyCredits(c *fiber.Ctx) error {
	res, err := crazycredits.CrazyCredits(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
