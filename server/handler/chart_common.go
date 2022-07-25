package handler

import (
	"github.com/Scrip7/imdb-api/pkg/chart/common"
	"github.com/gofiber/fiber/v2"
)

// ChartCommon returns list of titles based on the given chart type
// @Summary     Scrape common chart types
// @Description This endpoint can scrape common chart types that have the same response structure.
// @Description
// @Description Specifying the chart type based on the table below is required.
// @Description | URL | Page Title | Page Description |
// @Description | --- | --- | --- |
// @Description | `https://imdb.com/chart/top` | IMDb Top 250 Movies | IMDb Top 250 as rated by regular IMDb voters. |
// @Description | `https://imdb.com/chart/toptv` | Top Rated TV Shows | Top 250 as rated by IMDb Users |
// @Description | `https://imdb.com/chart/top-english-movies` | Top Rated English Movies | Top 250 English-language movies as rated by IMDb Users |
// @Description | `https://imdb.com/chart/bottom` | Lowest Rated Movies | Bottom 100 as voted by IMDb Users |
// @Tags        Chart
// @Param       type query    string true "Chart Type" Enums(top, toptv, top-english-movies, bottom)
// @Success     200  {object} pipe.ChartCommonTransform{}
// @Failure     500  {object} server.httpError
// @Router      /chart/common [get]
func ChartCommon(c *fiber.Ctx) error {
	res, err := common.Common(c.Query("type"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
