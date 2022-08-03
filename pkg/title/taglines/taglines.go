package taglines

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/title/taglines/parser"
	"github.com/Scrip7/imdb-api/pkg/title/taglines/pipe"
)

func Taglines(id string) (*pipe.TitleTaglinesTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_TAGLINES, id)
	res, err := client.Get(url, &client.GetOptions{})
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := parser.GetTaglines(doc)

	return transform, nil
}
