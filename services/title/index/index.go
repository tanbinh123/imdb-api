package index

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gosimple/slug"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/utils"
)

func Debug(id string) (*PageProps, error) {
	url := fmt.Sprintf(constants.TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := utils.ParseHTML(*res)
	if err != nil {
		return nil, err
	}

	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData TitleIndex
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}

	return &nextData.Props.PageProps, nil
}

func Index(id string) (*IndexTransform, error) {
	url := fmt.Sprintf(constants.TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := utils.ParseHTML(*res)
	if err != nil {
		return nil, err
	}

	scriptJSON := doc.Find("script[type=\"application/ld+json\"]").First()
	var data scriptResponse
	if err := json.Unmarshal([]byte(scriptJSON.Text()), &data); err != nil {
		return nil, err
	}

	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData TitleIndex
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}

	//
	// Begin Transformation
	//
	transform := IndexTransform{
		ID:   nextData.Props.PageProps.MainColumnData.ID,
		Type: slug.Make(nextData.Props.PageProps.MainColumnData.TitleType.ID),
		Title: title{
			Text:     nextData.Props.PageProps.MainColumnData.TitleText.Text,
			Original: nextData.Props.PageProps.MainColumnData.OriginalTitleText.Text,
			AKA:      getTitleAKA(nextData.Props.PageProps.MainColumnData.Akas.Edges),
		},
		Plot:            nextData.Props.PageProps.AboveTheFoldData.Plot.PlotText.PlainText,
		IsAdult:         nextData.Props.PageProps.MainColumnData.IsAdult,
		CanHaveEpisodes: nextData.Props.PageProps.MainColumnData.CanHaveEpisodes,
		Popularity: popularity{
			Rank:       nextData.Props.PageProps.AboveTheFoldData.MeterRanking.CurrentRank,
			Difference: getRankingDifference(nextData.Props.PageProps.AboveTheFoldData.MeterRanking.RankChange),
		},
		Images: images{
			Total: nextData.Props.PageProps.MainColumnData.TitleMainImages.Total,
			Primary: primaryImage{
				ID:  nextData.Props.PageProps.MainColumnData.PrimaryImage.ID,
				URL: data.Image,
			},
			Items: getImageItems(nextData.Props.PageProps.MainColumnData.TitleMainImages.Edges),
		},
		Videos: videos{
			Total: nextData.Props.PageProps.MainColumnData.Videos.Total,
			Items: getVideoItems(nextData.Props.PageProps.AboveTheFoldData.PrimaryVideos.Edges),
		},
		Reviews: reviews{
			Featured: getFeaturedReviews(nextData.Props.PageProps.MainColumnData.FeaturedReviews.Edges),
			Users: usersReviews{
				Total: nextData.Props.PageProps.MainColumnData.Reviews.Total,
			},
			External: externalReviews{
				Total: nextData.Props.PageProps.AboveTheFoldData.CriticReviewsTotal.Total,
			},
		},
		Keywords: keyword{
			Total: nextData.Props.PageProps.AboveTheFoldData.Keywords.Total,
			Items: strings.Split(data.Keywords, ","),
		},

		// TODO: episodes
	}

	return &transform, nil
}

func getTitleAKA(edges []AKAEdge) []string {
	var items []string

	for _, v := range edges {
		items = append(items, v.Node.Text)
	}

	return items
}

func getRankingDifference(input RankChange) int64 {
	if strings.ToLower(input.ChangeDirection) == "down" {
		return -(input.Difference)
	}
	return input.Difference
}

func getImageItems(edges []TitleMainImagesEdge) []imageItem {
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

func getVideoItems(edges []PrimaryVideosEdge) []videoItem {
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

func getVideoPlaybackItems(urls []URL) []playbackItem {
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
