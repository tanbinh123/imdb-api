package index

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/utils"
)

// TODO: Remove this method after implementation
func Debug(id string) (*pageProps, error) {
	url := fmt.Sprintf(constants.URL_TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := utils.ParseHTML(*res)
	if err != nil {
		return nil, err
	}

	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData titleIndex
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}

	return &nextData.Props.PageProps, nil
}

func Index(id string) (*IndexTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := utils.ParseHTML(*res)
	if err != nil {
		return nil, err
	}

	// Extract the first JSON string, which is usually on top of the web page's source code
	// this is manually generated based on "schema.org" schemas,
	// and it's unlikely to change in the future.
	scriptJSON := doc.Find("script[type=\"application/ld+json\"]").First()
	var data scriptResponse
	if err := json.Unmarshal([]byte(scriptJSON.Text()), &data); err != nil {
		return nil, err
	}

	// Extract the second JSON string
	// which is usually in the middle or bottom of the web page's source code
	// the Next.JS framework automatically generates this.
	// Most likely, the structure of this JSON object is going to change.
	// But as long as they use the Next.JS framework, the approach remains the same.
	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData titleIndex
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
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
	transform := IndexTransform{
		ID: main.ID,
		Validate: validate{
			Type:            titleType,
			IsMovie:         titleType == constants.TITLE_TYPE_MOVIE,
			IsSeries:        titleType == constants.TITLE_TYPE_SERIES || titleType == constants.TITLE_TYPE_EPISODE,
			IsEpisode:       titleType == constants.TITLE_TYPE_EPISODE,
			IsAdult:         main.IsAdult,
			CanHaveEpisodes: main.CanHaveEpisodes,
		},
		Title: title{
			Text:     main.TitleText.Text,
			Original: main.OriginalTitleText.Text,
			Slug:     slug.Make(main.OriginalTitleText.Text),
			AKA:      getTitleAKA(main.Akas.Edges),
		},
		Genres: getGenres(above.Genres.Genres),
		Plot:   above.Plot.PlotText.PlainText,
		Popularity: popularity{
			Rank:       above.MeterRanking.CurrentRank,
			Difference: getRankingDifference(above.MeterRanking.RankChange),
		},
		Images: images{
			Total: main.TitleMainImages.Total,
			Primary: primaryImage{
				ID:  main.PrimaryImage.ID,
				URL: data.Image,
			},
			Items: getImageItems(main.TitleMainImages.Edges),
		},
		Videos: videos{
			Total: main.Videos.Total,
			Items: getVideoItems(above.PrimaryVideos.Edges),
		},
		Reviews: reviews{
			Featured: getFeaturedReviews(main.FeaturedReviews.Edges),
			Users: usersReviews{
				Total: main.Reviews.Total,
			},
			External: externalReviews{
				Total: above.CriticReviewsTotal.Total,
			},
		},
		FAQ: faq{
			Total: main.FaqsTotal.Total,
			Items: getFaqItems(main.Faqs.Edges),
		},
		Trivia: trivia{
			Total: main.TriviaTotal.Total,
			Items: getTriviaItems(main.Trivia.Edges),
		},
		Keywords: keyword{
			Total: above.Keywords.Total,
			Items: strings.Split(data.Keywords, ","),
		},
		Series: series{
			ID: seriesID{
				Parent:          above.Series.Series.ID,
				EpisodeNext:     above.Series.NextEpisode.ID,
				EpisodePrevious: above.Series.PreviousEpisode.ID,
			},
			Title: seriesTitle{
				Text:     above.Series.Series.TitleText.Text,
				Original: above.Series.Series.OriginalTitleText.Text,
				Slug:     slug.Make(above.Series.Series.OriginalTitleText.Text),
			},
			Current: seriesCurrent{
				Episode: above.Series.EpisodeNumber.EpisodeNumber,
				Season:  above.Series.EpisodeNumber.SeasonNumber,
			},
			ReleaseYear: releaseYear{
				From: above.Series.Series.ReleaseYear.Year,
				To:   above.Series.Series.ReleaseYear.EndYear,
			},
		},
	}

	return &transform, nil
}

func getTitleAKA(edges []akaEdge) []string {
	var items []string

	for _, v := range edges {
		items = append(items, v.Node.Text)
	}

	return items
}

func getGenres(genres []withTextAndID) []genre {
	var items []genre

	for _, v := range genres {
		items = append(items, genre{
			Name: v.Text,
			Slug: slug.Make(v.ID),
		})
	}

	return items
}

func getRankingDifference(input rankChange) int64 {
	if strings.ToLower(input.ChangeDirection) == "down" {
		return -(input.Difference)
	}
	return input.Difference
}

func getImageItems(edges []imageEdge) []imageItem {
	var items []imageItem

	for _, v := range edges {
		items = append(items, imageItem{
			ID:      v.Node.ID,
			Width:   v.Node.Width,
			Height:  v.Node.Height,
			URL:     v.Node.URL,
			Caption: v.Node.Caption.PlainText,
		})
	}

	return items
}

func getVideoItems(edges []primaryVideosEdge) []videoItem {
	var items []videoItem

	for _, v := range edges {
		items = append(items, videoItem{
			ID:          v.Node.ID,
			Type:        slug.Make(v.Node.ContentType.DisplayName.Value),
			Title:       v.Node.Name.Value,
			Description: v.Node.Description.Value,
			Duration:    v.Node.Runtime.Value,
			Thumbnail: thumbnail{
				URL:    v.Node.Thumbnail.URL,
				Width:  v.Node.Thumbnail.Width,
				Height: v.Node.Thumbnail.Height,
			},
			Playback: getVideoPlaybackItems(v.Node.PlaybackURLs),
			IsMature: v.Node.IsMature,
		})
	}

	return items
}

func getVideoPlaybackItems(urls []urlWrapper) []playbackItem {
	var items []playbackItem

	for _, v := range urls {
		items = append(items, playbackItem{
			Quality:  v.DisplayName.Value,
			MIMEType: v.MIMEType,
			URL:      v.URL,
		})
	}

	return items
}

func getFeaturedReviews(edges []featuredReviewEdge) []featuredReviewItem {
	var items []featuredReviewItem

	for _, v := range edges {
		items = append(items, featuredReviewItem{
			ID: v.Node.ID,
			Author: reviewAuthor{
				ID:       v.Node.Author.UserID,
				Nickname: v.Node.Author.NickName,
				Rating:   v.Node.AuthorRating,
			},
			Summary:   v.Node.Summary.OriginalText,
			Text:      utils.ParseHTMLToString(v.Node.Text.OriginalText.PlaidHTML),
			Likes:     v.Node.Helpfulness.UpVotes,
			Dislikes:  v.Node.Helpfulness.DownVotes,
			CreatedAt: v.Node.SubmissionDate,
		})
	}

	return items
}

func getFaqItems(edges []faqEdge) []faqItem {
	var items []faqItem

	for _, v := range edges {
		items = append(items, faqItem{
			ID:       v.Node.ID,
			Question: v.Node.Question.PlainText,
		})
	}

	return items
}

func getTriviaItems(edges []triviaEdge) []string {
	var items []string

	for _, v := range edges {
		items = append(items, utils.ParseHTMLToString(v.Node.Text.PlaidHTML))
	}

	return items
}
