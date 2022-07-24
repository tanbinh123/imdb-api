package client_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/client"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	url := "https://m.imdb.com/title/tt20256528"
	res, err := client.Get(url)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
