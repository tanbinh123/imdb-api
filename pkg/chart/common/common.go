package common

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/pkg/chart/common/parser"
	"github.com/Scrip7/imdb-api/pkg/chart/common/pipe"
)

type CommonChartType = string

const (
	TYPE_TV                 CommonChartType = "tv"
	TYPE_TOP_TV             CommonChartType = "toptv"
	TYPE_TOP_ENGLISH_MOVIES CommonChartType = "top-english-movies"
	TYPE_BOTTOM             CommonChartType = "bottom"
)

func Common(t CommonChartType) (*pipe.ChartCommonTransform, error) {
	url := fmt.Sprintf(constants.URL_CHART_COMMON, t)
	res, err := client.Get(url, &client.GetOptions{})
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	transform := &pipe.ChartCommonTransform{
		Type:   t,
		Titles: parser.GetChartData(doc),
	}

	return transform, nil
}
