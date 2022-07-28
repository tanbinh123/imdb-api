package moviemeter

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/chart/moviemeter/parser"
	"github.com/Scrip7/imdb-api/pkg/chart/moviemeter/pipe"
)

func MovieMeter() (*pipe.ChartMovieMeterTransform, error) {
	res, err := client.Get(constants.URL_CHART_MOVIE_METER, &client.GetOptions{})
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := &pipe.ChartMovieMeterTransform{
		Titles: parser.GetChartData(doc),
	}

	return transform, nil
}
