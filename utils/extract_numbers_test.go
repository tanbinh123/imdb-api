package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractNumbers(t *testing.T) {
	res := ExtractNumbers("abc123def")
	assert.Equal(t, res, "123")
}
