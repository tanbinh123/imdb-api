package crazycredits_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/title/crazycredits"
	"github.com/stretchr/testify/assert"
)

func TestCrazyCredits(t *testing.T) {
	res, err := crazycredits.CrazyCredits("tt0108778") // Friends
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.GreaterOrEqual(t, len(res.Items), 2) // Items should be an empty array
}
