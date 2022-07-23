package title

import (
	"github.com/Scrip7/imdb-api/services/title/index"
	"github.com/gofiber/fiber/v2"
)

// Index is a function to get information about IMDb title
// @Summary     Get IMDb title by ID
// @Description The ID should start with "tt" at the beginning.
// @Description | Title Text | IMDb ID |
// @Description | --- | --- |
// @Description | Spider-Man: No Way Home | `tt10872600` |
// @Description | Spider-Man: Far from Home | `tt6320628` |
// @Description | Spider-Man: Homecoming | `tt2250912` |
// @Description | Avengers: Endgame | `tt4154796` |
// @Description | Avengers: Infinity War | `tt4154756` |
// @Description | The Dark Knight | `tt0468569` |
// @Description | The Godfather | `tt0068646` |
// @Description | Friends | `tt0108778` |
// @Description | Breaking Bad | `tt0903747` |
// @Description | Chernobyl | `tt7366338` |
// @Tags        Title
// @Param       id  path     string true "Title ID"
// @Success     200 {object} pipe.IndexTransform{}
// @Failure     404 {object} server.HTTPError
// @Failure     500 {object} server.HTTPError
// @Router      /title/{id} [get]
func Index(c *fiber.Ctx) error {
	res, err := index.Index(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(res)
}
