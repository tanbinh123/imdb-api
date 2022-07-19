package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovie(t *testing.T) {
	res, err := Index("tt20256528") // What Is a Woman?
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestTVSeries(t *testing.T) {
	res, err := Index("tt0108778") // Friends
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestTVSeriesEpisode(t *testing.T) {
	res, err := Index("tt0583436") // Friends S10.E1
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
