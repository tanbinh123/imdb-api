package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	res, err := Index("tt20256528")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
