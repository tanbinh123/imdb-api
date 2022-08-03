package crazycredits

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/title/crazycredits/parser"
	"github.com/Scrip7/imdb-api/pkg/title/crazycredits/pipe"
)

func CrazyCredits(id string) (*pipe.TitleCrazyCreditsTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_CRAZY_CREDITS, id)
	res, err := client.Get(url, &client.GetOptions{})
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := &pipe.TitleCrazyCreditsTransform{
		Items: parser.GetItems(doc),
	}

	return transform, nil
}
