package handler

import (
	"github.com/Scrip7/imdb-api/pkg/title/keywords"
	"github.com/gofiber/fiber/v2"
)

// TitleKeywords is a function that returns list of titles keywords
// @Summary Get titles keywords
// @Tags    Title
// @Param   id  path     string true "Title ID"
// @Success 200 {object} pipe.TitleKeywordsTransform{}
// @Failure 404 {object} server.httpError
// @Failure 500 {object} server.httpError
// @Router  /title/{id}/keywords [get]
func TitleKeywords(c *fiber.Ctx) error {
	res, err := keywords.Keywords(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
