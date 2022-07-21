package index

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
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

	doc, err := goquery.NewDocumentFromReader(*res)
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

	// Values are refactored with the Extract Method technique
	// https://refactoring.guru/extract-method
	transform := pipe.IndexTransform{
		ID: main.ID,
		Validate: pipe.Validate{
			Type:            titleType,
			IsMovie:         titleType == constants.TITLE_TYPE_MOVIE,
			IsSeries:        titleType == constants.TITLE_TYPE_SERIES || titleType == constants.TITLE_TYPE_EPISODE,
			IsEpisode:       titleType == constants.TITLE_TYPE_EPISODE,
			IsAdult:         main.IsAdult,
			CanHaveEpisodes: main.CanHaveEpisodes,
		},
		Title: pipe.Title{
			Text:     main.TitleText.Text,
			Original: main.OriginalTitleText.Text,
			Slug:     slug.Make(main.OriginalTitleText.Text),
			AKA:      parser.GetTitleAKA(main.AKAs.Edges),
		},
		Genres: parser.GetGenres(above.Genres.Genres),
		Plot:   above.Plot.PlotText.PlainText,
		Popularity: pipe.Popularity{
			Rank:       above.MeterRanking.CurrentRank,
			Difference: parser.GetRankingDifference(above.MeterRanking.RankChange),
		},
		Images: pipe.Images{
			Total: main.TitleMainImages.Total,
			Primary: pipe.PrimaryImage{
				ID: above.PrimaryImage.ID,
				// URL:        schemaData.Image,
				URL:        above.PrimaryImage.URL,
				Width:      above.PrimaryImage.Width,
				Height:     above.PrimaryImage.Height,
				Caption:    above.PrimaryImage.Caption.PlainText,
				Thumbnails: parser.GetPrimaryImageThumbnails(doc),
			},
			Items: parser.GetImageItems(main.TitleMainImages.Edges),
		},
		Videos: pipe.Videos{
			Total:     main.Videos.Total,
			Primaries: parser.GetPrimaryVideos(above.PrimaryVideos.Edges),
			Items:     parser.GetVideoItems(main.VideoStrip.Edges),
		},
		Reviews: pipe.Reviews{
			Featured: parser.GetFeaturedReviews(main.FeaturedReviews.Edges),
			Users: pipe.UsersReviews{
				Total: main.Reviews.Total,
			},
			External: pipe.ExternalReviews{
				Total: above.CriticReviewsTotal.Total,
			},
		},
		FAQ: pipe.FAQ{
			Total: main.FaqsTotal.Total,
			Items: parser.GetFAQItems(main.Faqs.Edges),
		},
		Trivia: pipe.Trivia{
			Total: main.TriviaTotal.Total,
			Items: parser.GetTriviaItems(main.Trivia.Edges),
		},
		Keywords: pipe.Keywords{
			Total: above.Keywords.Total,
			Items: parser.GetKeywords(above.Keywords.Edges),
		},
		Series: pipe.Series{
			ID: pipe.SeriesID{
				Parent:          above.Series.Series.ID,
				EpisodeNext:     above.Series.NextEpisode.ID,
				EpisodePrevious: above.Series.PreviousEpisode.ID,
			},
			Title: pipe.SeriesTitle{
				Text:     above.Series.Series.TitleText.Text,
				Original: above.Series.Series.OriginalTitleText.Text,
				Slug:     slug.Make(above.Series.Series.OriginalTitleText.Text),
			},
			Current: pipe.SeriesCurrent{
				Episode: above.Series.EpisodeNumber.EpisodeNumber,
				Season:  above.Series.EpisodeNumber.SeasonNumber,
			},
			ReleaseYear: pipe.ReleaseYear{
				From: above.Series.Series.ReleaseYear.Year,
				To:   above.Series.Series.ReleaseYear.EndYear,
			},
		},
		Soundtracks: parser.GetSoundtracks(main.Soundtrack.Edges),
		Related:     parser.GetRelatedTitles(main.MoreLikeThisTitles.Edges),
	}

	return &transform, nil
}
