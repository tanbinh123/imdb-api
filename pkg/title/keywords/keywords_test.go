package keywords_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/title/keywords"
	"github.com/stretchr/testify/assert"
)

func TestKeywords(t *testing.T) {
  res, err := keywords.Keywords("tt0108778") // Friends
  assert.Nil(t, err)
  assert.NotNil(t, res)
}