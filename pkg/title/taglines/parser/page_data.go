package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/title/taglines/pipe"
)

func GetTaglines(doc *goquery.Document) *pipe.TitleTaglinesTransform {
	items := pipe.TitleTaglinesTransform{}

	doc.Find("div#main div.article div#taglines_content div.soda").Each(func(i int, e *goquery.Selection) {
		text := strings.TrimSpace(e.Text())
		if text == "" {
			return
		}

		items = append(items, text)
	})

	return &items
}
