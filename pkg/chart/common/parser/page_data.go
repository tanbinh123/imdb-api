package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/chart/common/pipe"
	"github.com/Scrip7/imdb-api/utils"
	"github.com/gosimple/slug"
)

func GetChartData(doc *goquery.Document) []*pipe.ChatItem {
	var items []*pipe.ChatItem

	doc.Find("section[id=chart-content] div.media").Each(func(i int, s *goquery.Selection) {
		// Extract title ID (starts with "tt")
		ID, exists := s.Find("span").First().Attr("data-tconst")
		if !exists {
			return
		}

		lines := strings.Split(
			strings.TrimSpace(s.Find("span.media-body h4").Text()), // "1.\nTEXT\n(2022)"
			"\n")
		if len(lines) != 3 {
			return
		}

		// Release year
		releaseYear, err := strconv.ParseFloat(utils.ExtractNumbers(lines[2]), 64)
		if err != nil {
			return
		}

		// Extract poster thumbnail image URL
		thumb, exists := s.Find("img").First().Attr("src")
		if !exists {
			return
		}

		// Rating
		rating, err := strconv.ParseFloat(s.Find("span.media-body p span.imdb-rating").Text(), 64)
		if err != nil {
			return
		}

		items = append(items, &pipe.ChatItem{
			ID: ID,
			Title: &pipe.ChartItemTitle{
				Text: lines[1],
				Slug: slug.Make(lines[1]),
			},
			Thumbnail:   thumb,
			Rating:      rating,
			ReleaseYear: releaseYear,
		})
	})

	return items
}
