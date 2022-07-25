package common_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/pkg/chart/common"
	"github.com/stretchr/testify/assert"
)

func TestCommonTV(t *testing.T) {
	res, err := common.Common(common.TYPE_TV)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestCommonTopTV(t *testing.T) {
	res, err := common.Common(common.TYPE_TOP_TV)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestCommonTopEnglishMovies(t *testing.T) {
	res, err := common.Common(common.TYPE_TOP_ENGLISH_MOVIES)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestCommonBottom(t *testing.T) {
	res, err := common.Common(common.TYPE_BOTTOM)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
