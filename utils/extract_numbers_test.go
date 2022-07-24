package utils_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/utils"
	"github.com/stretchr/testify/assert"
)

func TestExtractNumbers(t *testing.T) {
	res := utils.ExtractNumbers("abc123def")
	assert.Equal(t, res, "123")
}
