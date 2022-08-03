package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/title/keywords/pipe"
	"github.com/gosimple/slug"
)

func GetKeywords(doc *goquery.Document) []*pipe.TitleKeyword {
	var items []*pipe.TitleKeyword

	doc.Find("div#main table tbody tr td").Each(func(i int, td *goquery.Selection) {
		keyword, exists := td.Attr("data-item-keyword")
		if !exists {
			return
		}

		ID, exists := td.Find("span.interesting-cast-vote").Attr("data-item-id")
		if !exists {
			return
		}

		items = append(items, &pipe.TitleKeyword{
			ID:   ID,
			Text: keyword,
			Slug: slug.Make(keyword),
		})
	})

	return items
}
