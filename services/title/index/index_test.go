package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovie(t *testing.T) {
	id := "tt20256528" // What Is a Woman?

	res, err := Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "What Is a Woman?", res.Title.Original)
}

func TestTVSeries(t *testing.T) {
	id := "tt0108778" // Friends

	res, err := Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Friends", res.Title.Original)
}

func TestTVSeriesEpisode(t *testing.T) {
	id := "tt0583436" // Friends S10.E1 (The One After Joey and Rachel Kiss)

	res, err := Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "The One After Joey and Rachel Kiss", res.Title.Original)
}

func TestVideoGame(t *testing.T) {
	id := "tt2140553" // The Last of Us

	res, err := Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "The Last of Us", res.Title.Original)
}
