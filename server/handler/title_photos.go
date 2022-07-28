package handler

import (
	"strconv"

	"github.com/Scrip7/imdb-api/pkg/title/photos"
	"github.com/gofiber/fiber/v2"
)

// TitlePhotos is a function that returns list of titles photos with pagination
// @Summary Titles photos
// @Tags    Title
// @Param   id   path     string true "Title ID"
// @Param   page query    string true "Page Number" Format(integer) Default(1)
// @Success 200  {object} pipe.TitlePhotosTransform{}
// @Failure 404  {object} server.httpError
// @Failure 500  {object} server.httpError
// @Router  /title/{id}/photos [get]
func TitlePhotos(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return err
	}

	res, err := photos.Photos(c.Params("id"), page)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
