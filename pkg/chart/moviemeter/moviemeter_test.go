package moviemeter_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/chart/moviemeter"
	"github.com/stretchr/testify/assert"
)

func TestMovieMeter(t *testing.T) {
	res, err := moviemeter.MovieMeter()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
