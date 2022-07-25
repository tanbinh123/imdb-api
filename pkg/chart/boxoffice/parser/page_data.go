package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/pkg/chart/boxoffice/pipe"
	"github.com/gosimple/slug"
)

func GetChartData(doc *goquery.Document) []*pipe.BoxOfficeItem {
	var items []*pipe.BoxOfficeItem

	doc.Find("section[id=chart-content] div.media").Each(func(i int, s *goquery.Selection) {
		// Extract title ID (starts with "tt")
		ID, exists := s.Find("span").First().Attr("data-tconst")
		if !exists {
			return
		}

		text := strings.TrimSpace(s.Find("span.media-body h4").Text())

		// Extract poster thumbnail image URL
		thumb, exists := s.Find("img").First().Attr("src")
		if !exists {
			return
		}

		descriptionItems := strings.Split(strings.TrimSpace(
			s.Find("span.media-body p").First().Text(), // Weekend Gross: $46.6M\nTotal Gross: $233.9M\nWeeks Since Release: 2
		), "\n\n")

		items = append(items, &pipe.BoxOfficeItem{
			ID: ID,
			Title: &pipe.BoxOfficeItemTitle{
				Text: text,
				Slug: slug.Make(text),
			},
			Thumbnail:         thumb,
			WeekendGross:      getWeekendGross(&descriptionItems),
			TotalGross:        getTotalGross(&descriptionItems),
			WeeksSinceRelease: getWeeksSinceRelease(&descriptionItems),
		})
	})

	return items
}

func getWeekendGross(desc *[]string) string {
	for _, v := range *desc {
		if strings.HasPrefix(v, "Weekend Gross: ") {
			return strings.Replace(v, "Weekend Gross: ", "", 1)
		}
	}

	return ""
}

func getTotalGross(desc *[]string) string {
	for _, v := range *desc {
		if strings.HasPrefix(v, "Total Gross: ") {
			return strings.Replace(v, "Total Gross: ", "", 1)
		}
	}

	return ""
}

func getWeeksSinceRelease(desc *[]string) string {
	for _, v := range *desc {
		if strings.HasPrefix(v, "Weeks Since Release: ") {
			return strings.Replace(v, "Weeks Since Release: ", "", 1)
		}
	}

	return ""
}
