package taglines_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/title/taglines"
	"github.com/stretchr/testify/assert"
)

func TestTaglines(t *testing.T) {
	res, err := taglines.Taglines("tt0108778") // Friends
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.GreaterOrEqual(t, len(*res), 2) // Should not be an empty array
}
