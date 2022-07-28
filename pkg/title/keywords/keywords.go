package keywords

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/title/keywords/parser"
	"github.com/Scrip7/imdb-api/pkg/title/keywords/pipe"
)

func Keywords(id string) (*pipe.TitleKeywordsTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_KEYWORDS, id)
	res, err := client.Get(url, &client.GetOptions{})
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := &pipe.TitleKeywordsTransform{
		Keywords: parser.GetKeywords(doc),
	}

	return transform, nil
}
