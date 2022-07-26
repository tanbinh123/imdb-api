package handler

import (
	"github.com/Scrip7/imdb-api/pkg/chart/moviemeter"
	"github.com/gofiber/fiber/v2"
)

// ChartMovieMeter returns most popular movies as determined by IMDb Users
// @Summary Most Popular Movies
// @Tags    Chart
// @Success 200 {object} pipe.ChartMovieMeterTransform{}
// @Failure 500 {object} server.httpError
// @Router  /chart/moviemeter [get]
func ChartMovieMeter(c *fiber.Ctx) error {
	res, err := moviemeter.MovieMeter()
	if err != nil {
		return err
	}

	return c.JSON(res)
}
