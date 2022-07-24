package boxoffice_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/chart/boxoffice"
	"github.com/stretchr/testify/assert"
)

func TestBoxOffice(t *testing.T) {
	res, err := boxoffice.BoxOffice()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
