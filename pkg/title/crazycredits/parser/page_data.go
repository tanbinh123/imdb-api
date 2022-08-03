package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/title/crazycredits/pipe"
)

func GetItems(doc *goquery.Document) []*pipe.CrazyCreditItem {
	items := []*pipe.CrazyCreditItem{}

	doc.Find("div.container div.list div.row").Each(func(i int, row *goquery.Selection) {
		ID, exists := row.Attr("data-cz-item")
		if !exists {
			return
		}

		text := strings.TrimSpace(row.Find("p").Text())
		if text == "" {
			return
		}

		items = append(items, &pipe.CrazyCreditItem{
			ID:   ID,
			Text: text,
		})
	})

	return items
}
