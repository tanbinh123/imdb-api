package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	url := "https://m.imdb.com/title/tt20256528"
	res, err := Get(url)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
