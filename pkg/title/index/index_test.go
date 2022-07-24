package index_test

import (
	"testing"

	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/title/index"
	"github.com/stretchr/testify/assert"
)

func TestMovie(t *testing.T) {
	id := "tt20256528" // What Is a Woman?

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "What Is a Woman?", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_MOVIE)
}

func TestTVSeries(t *testing.T) {
	id := "tt0108778" // Friends

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Friends", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_TV_SERIES)
}

func TestTVSeriesEpisode(t *testing.T) {
	id := "tt0583436" // Friends S10.E1 (The One After Joey and Rachel Kiss)

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "The One After Joey and Rachel Kiss", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_TV_EPISODE)
}

func TestTVShort(t *testing.T) {
	id := "tt4448662" // Mr Bean: Funeral

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Mr Bean: Funeral", res.Title.Text)
	assert.Equal(t, "Funeral", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_TV_SHORT)
}

func TestVideoGame(t *testing.T) {
	id := "tt2140553" // The Last of Us

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "The Last of Us", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_VIDEO_GAME)
}

func TestTVSpecial(t *testing.T) {
	id := "tt21106500" // Bill Burr: Live at Red Rocks

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Bill Burr: Live at Red Rocks", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_TV_SPECIAL)
}

func TestTVMiniSeries(t *testing.T) {
	id := "tt7366338" // Chernobyl

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Chernobyl", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_TV_MINI_SERIES)
}

func TestVideo(t *testing.T) {
	id := "tt18258584" // Jackass 4.5

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Jackass 4.5", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_VIDEO)
}

func TestMusicVideo(t *testing.T) {
	id := "tt4272746" // DJ Snake and Lil Jon: Turn Down for What

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "DJ Snake and Lil Jon: Turn Down for What", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_MUSIC_VIDEO)
}

func TestPodcastSeries(t *testing.T) {
	id := "tt12326830" // The Sandman

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "The Sandman", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_PODCAST_SERIES)
}

func TestPodcastEpisode(t *testing.T) {
	id := "tt13998532" // Live in Toronto!

	res, err := index.Index(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, id, res.ID)
	assert.Equal(t, "Live in Toronto!", res.Title.Original)
	assert.Equal(t, res.Validate.Type, constants.TITLE_TYPE_PODCAST_EPISODE)
}
