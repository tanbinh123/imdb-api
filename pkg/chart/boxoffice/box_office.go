package boxoffice

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/chart/boxoffice/parser"
	"github.com/Scrip7/imdb-api/pkg/chart/boxoffice/pipe"
)

func BoxOffice() (*pipe.BoxOfficeTransform, error) {
	res, err := client.Get(constants.URL_CHART_BOX_OFFICE)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := &pipe.BoxOfficeTransform{
		Titles: parser.GetChartData(doc),
	}

	return transform, nil
}
