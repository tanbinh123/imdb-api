package handler

import (
	"github.com/Scrip7/imdb-api/pkg/title/taglines"
	"github.com/gofiber/fiber/v2"
)

// TitleTaglines is a function that returns list of titles taglines
// @Summary Get titles taglines
// @Tags    Title
// @Param   id  path     string true "Title ID"
// @Success 200 {object} pipe.TitleTaglinesTransform{}
// @Failure 404 {object} server.httpError
// @Failure 500 {object} server.httpError
// @Router  /title/{id}/taglines [get]
func TitleTaglines(c *fiber.Ctx) error {
	res, err := taglines.Taglines(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
