package photos

import (
	"fmt"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/title/photos/parser"
	"github.com/Scrip7/imdb-api/pkg/title/photos/pipe"
)

func Photos(id string, page int) (*pipe.TitlePhotosTransform, error) {
	if page <= 0 {
		page = 1
	}

	url := fmt.Sprintf(constants.URL_TITLE_PHOTOS, id)
	res, err := client.Get(url, &client.GetOptions{
		Params: &client.Map{
			"page":          fmt.Sprint(page),
			"retina":        "true",
			"maximagewidth": "2048",
		},
	})
	if err != nil {
		return nil, err
	}

	resObj, err := parser.JSONDecode(res)
	if err != nil {
		return nil, err
	}

	transform := &pipe.TitlePhotosTransform{
		Photos: parser.GetPhotos(resObj.Data),
		Meta:   parser.GetMetaData(resObj.Meta),
	}

	return transform, nil
}
