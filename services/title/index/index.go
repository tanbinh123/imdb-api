package index

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/services/title/index/parser"
	"github.com/Scrip7/imdb-api/services/title/index/pipe"
)

func Index(id string) (*pipe.IndexTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*res))
	if err != nil {
		return nil, err
	}

	// schemaData, err := parser.GetSchemaOrgData(doc)
	// if err != nil {
	// 	return nil, err
	// }

	nextData, err := parser.GetNextJSData(doc)
	if err != nil {
		return nil, err
	}

	//
	// Begin Transformation
	//

	// To improve code readability
	above := &nextData.Props.PageProps.AboveTheFoldData
	main := &nextData.Props.PageProps.MainColumnData

	// The lowerCamelCase is the structure we follow based on constants
	// to provide a consistent stable API
	titleType := strcase.ToLowerCamel(main.TitleType.ID)

	validate := &pipe.Validate{
		Type:            titleType,
		IsMovie:         titleType == constants.TITLE_TYPE_MOVIE,
		IsSeries:        titleType == constants.TITLE_TYPE_TV_SERIES,
		IsEpisode:       titleType == constants.TITLE_TYPE_TV_EPISODE,
		IsAdult:         main.IsAdult,
		CanHaveEpisodes: main.CanHaveEpisodes,
	}

	// Values are refactored with the Extract Method technique
	// https://refactoring.guru/extract-method
	transform := &pipe.IndexTransform{
		ID:          main.ID,
		Validate:    validate,
		Title:       parser.GetTitle(main.TitleText.Text, main.OriginalTitleText.Text, main.AKAs.Edges),
		Genres:      parser.GetGenres(above.Genres.Genres),
		Plot:        above.Plot.PlotText.PlainText,
		Popularity:  parser.GetPopularity(above.MeterRanking),
		Images:      parser.GetImages(main.TitleMainImages.Total, above.PrimaryImage, main.TitleMainImages.Edges, doc),
		Videos:      parser.GetVideos(main.Videos.Total, above.PrimaryVideos.Edges, main.VideoStrip.Edges),
		Cast:        parser.GetCast(main.Cast.Edges),
		Reviews:     parser.GetReviews(main.Reviews.Total, above.CriticReviewsTotal.Total, main.FeaturedReviews.Edges),
		FAQ:         parser.GetFAQ(main.FaqsTotal.Total, main.Faqs.Edges),
		Trivia:      parser.GetTrivia(main.TriviaTotal.Total, main.Trivia.Edges),
		Keywords:    parser.GetKeywords(above.Keywords),
		Series:      parser.GetSeries(above.Series),
		Soundtracks: parser.GetSoundtracks(main.Soundtrack.Edges),
		Related:     parser.GetRelatedTitles(main.MoreLikeThisTitles.Edges),
	}

	return transform, nil
}
