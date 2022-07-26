package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/chart/moviemeter/pipe"
	"github.com/Scrip7/imdb-api/utils"
	"github.com/gosimple/slug"
)

func GetChartData(doc *goquery.Document) []*pipe.ChartMovieMeterItem {
	var items []*pipe.ChartMovieMeterItem

	doc.Find("section[id=chart-content] div.media").Each(func(i int, s *goquery.Selection) {
		// Extract title ID (starts with "tt")
		ID, exists := s.Find("span").First().Attr("data-tconst")
		if !exists {
			return
		}

		lines := strings.Split(
			strings.TrimSpace(s.Find("span.media-body h4").Text()), // "1.\nTEXT\n(2022)"
			"\n")
		if len(lines) != 2 {
			return
		}

		// Release year
		releaseYear, err := strconv.ParseFloat(utils.ExtractNumbers(lines[1]), 64)
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

		items = append(items, &pipe.ChartMovieMeterItem{
			ID: ID,
			Title: &pipe.ChartMovieMeterItemTitle{
				Text: lines[0],
				Slug: slug.Make(lines[0]),
			},
			Thumbnail:     thumb,
			Rating:        rating,
			RankingChange: 0, // TODO: Pull Requests are welcome! "<span class="secondaryInfo">(<span class="up down"></span>9999)</span>"
			ReleaseYear:   releaseYear,
		})
	})

	return items
}
