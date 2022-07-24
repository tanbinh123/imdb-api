package handler

import (
	"github.com/Scrip7/imdb-api/pkg/chart/boxoffice"
	"github.com/gofiber/fiber/v2"
)

// ChartBoxOffice returns titles that are top of the box office chart
// @Summary Top Box Office (US)
// @Tags    Chart
// @Success 200 {object} pipe.BoxOfficeTransform{}
// @Failure 500 {object} server.httpError
// @Router  /chart/box-office [get]
func ChartBoxOffice(c *fiber.Ctx) error {
	res, err := boxoffice.BoxOffice()
	if err != nil {
		return err
	}

	return c.JSON(res)
}
