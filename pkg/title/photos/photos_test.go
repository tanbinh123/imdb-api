package photos_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/title/photos"
	"github.com/stretchr/testify/assert"
)

func TestPhotos(t *testing.T) {
	id := "tt0108778" // Friends

	res, err := photos.Photos(id, 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
