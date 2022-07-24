package boxoffice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxOffice(t *testing.T) {
	res, err := BoxOffice()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
