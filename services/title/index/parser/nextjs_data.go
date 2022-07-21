package parser

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/services/title/index/pipe"
	"github.com/Scrip7/imdb-api/utils"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
	"github.com/lukasbob/srcset"
)

// GetNextJSData extracts the second JSON string
// which is usually in the middle or bottom of the web page's source code
// the Next.JS framework automatically generates this.
// Most likely, the structure of this JSON object is going to change.
// But as long as they use the Next.JS framework, the approach remains the same.
func GetNextJSData(doc *goquery.Document) (*nextJSData, error) {
	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData nextJSData
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}
	return &nextData, nil
}

func GetTitleAKA(edges []akaEdge) []string {
	items := []string{}

	for _, v := range edges {
		items = append(items, v.Node.Text)
	}

	return items
}

func GetGenres(genres []withTextAndID) []pipe.Genre {
	items := []pipe.Genre{}

	for _, v := range genres {
		items = append(items, pipe.Genre{
			Name: v.Text,
			Slug: slug.Make(v.ID),
		})
	}

	return items
}

func GetRankingDifference(input rankChange) int64 {
	if strings.ToLower(input.ChangeDirection) == "down" {
		return -(input.Difference)
	}
	return input.Difference
}

func GetPrimaryImageThumbnails(doc *goquery.Document) []pipe.PrimaryImageThumbnail {
	items := []pipe.PrimaryImageThumbnail{}

	poster := doc.Find("section.ipc-page-section div[data-testid=\"hero-media__poster\"] img.ipc-image").First()
	srcsetRaw, exists := poster.Attr("srcset")
	if !exists {
		return items
	}
	srcsetSource := srcset.Parse(srcsetRaw)
	for _, v := range srcsetSource {
		items = append(items, pipe.PrimaryImageThumbnail{
			URL:   v.URL,
			Width: *v.Width,
			// IMDb does not return the "height" or "density"
		})
	}

	return items
}

func GetImageItems(edges []imageEdge) []pipe.ImageItem {
	items := []pipe.ImageItem{}

	for _, v := range edges {
		items = append(items, pipe.ImageItem{
			ID:      v.Node.ID,
			Width:   v.Node.Width,
			Height:  v.Node.Height,
			URL:     v.Node.URL,
			Caption: v.Node.Caption.PlainText,
		})
	}

	return items
}

func GetPrimaryVideos(edges []primaryVideosEdge) []pipe.VideoItemPrimary {
	items := []pipe.VideoItemPrimary{}

	for _, v := range edges {
		items = append(items, pipe.VideoItemPrimary{
			ID: v.Node.ID,
			Type: pipe.VideoTypeWrapper{
				Text: v.Node.ContentType.DisplayName.Value,
				Slug: slug.Make(v.Node.ContentType.DisplayName.Value),
			},
			Title:       v.Node.Name.Value,
			Description: v.Node.Description.Value,
			Duration:    v.Node.Runtime.Value,
			Thumbnail: pipe.Thumbnail{
				URL:    v.Node.Thumbnail.URL,
				Width:  v.Node.Thumbnail.Width,
				Height: v.Node.Thumbnail.Height,
			},
			Playback: GetVideoPlaybackItems(v.Node.PlaybackURLs),
			IsMature: v.Node.IsMature,
		})
	}

	return items
}

func GetVideoItems(edges []videoStripEdge) []pipe.VideoItem {
	items := []pipe.VideoItem{}

	for _, v := range edges {
		items = append(items, pipe.VideoItem{
			ID: v.Node.ID,
			Type: pipe.VideoTypeWrapper{
				Text: v.Node.ContentType.DisplayName.Value,
				Slug: slug.Make(v.Node.ContentType.DisplayName.Value),
			},
			Title:    v.Node.Name.Value,
			Duration: v.Node.Runtime.Value,
			Thumbnail: pipe.Thumbnail{
				URL:    v.Node.Thumbnail.URL,
				Width:  v.Node.Thumbnail.Width,
				Height: v.Node.Thumbnail.Height,
			},
		})
	}

	return items
}

func GetVideoPlaybackItems(urls []urlWrapper) []pipe.PlaybackItem {
	items := []pipe.PlaybackItem{}

	for _, v := range urls {
		items = append(items, pipe.PlaybackItem{
			Quality:  v.DisplayName.Value,
			MIMEType: v.MIMEType,
			URL:      v.URL,
		})
	}

	return items
}

func GetFeaturedReviews(edges []featuredReviewEdge) []pipe.FeaturedReviewItem {
	items := []pipe.FeaturedReviewItem{}

	for _, v := range edges {
		text := utils.ParseHTMLToPlainText(v.Node.Text.OriginalText.PlaidHTML)
		if text == "" {
			continue
		}
		items = append(items, pipe.FeaturedReviewItem{
			ID: v.Node.ID,
			Author: pipe.ReviewAuthor{
				ID:       v.Node.Author.UserID,
				Nickname: v.Node.Author.NickName,
				Rating:   v.Node.AuthorRating,
			},
			Summary:   v.Node.Summary.OriginalText,
			Text:      text,
			Likes:     v.Node.Helpfulness.UpVotes,
			Dislikes:  v.Node.Helpfulness.DownVotes,
			CreatedAt: v.Node.SubmissionDate,
		})
	}

	return items
}

func GetFAQItems(edges []faqEdge) []pipe.FAQItem {
	items := []pipe.FAQItem{}

	for _, v := range edges {
		items = append(items, pipe.FAQItem{
			ID:       v.Node.ID,
			Question: v.Node.Question.PlainText,
		})
	}

	return items
}

func GetTriviaItems(edges []triviaEdge) []string {
	items := []string{}

	for _, v := range edges {
		text := utils.ParseHTMLToPlainText(v.Node.Text.PlaidHTML)
		if text == "" {
			continue
		}
		items = append(items, text)
	}

	return items
}

func GetKeywords(edges []keywordEdge) []string {
	items := []string{}

	for _, v := range edges {
		items = append(items, v.Node.Text)
	}

	return items
}

func GetRelatedTitles(edges []moreLikeThisTitlesEdge) []pipe.RelatedTitle {
	items := []pipe.RelatedTitle{}

	for _, v := range edges {
		items = append(items, pipe.RelatedTitle{
			ID: v.Node.ID,
			Title: pipe.RelatedTitleName{
				Text:     v.Node.TitleText.Text,
				Original: v.Node.OriginalTitleText.Text,
				Slug:     slug.Make(v.Node.OriginalTitleText.Text),
			},
			Type:            strcase.ToLowerCamel(v.Node.TitleType.ID),
			CanHaveEpisodes: v.Node.CanHaveEpisodes,
			Poster: pipe.RelatedTitlePoster{
				ID:     v.Node.PrimaryImage.ID,
				Width:  v.Node.PrimaryImage.Width,
				Height: v.Node.PrimaryImage.Height,
				URL:    v.Node.PrimaryImage.URL,
			},
			ReleaseYear: pipe.ReleaseYear{
				From: v.Node.ReleaseYear.Year,
				To:   v.Node.ReleaseYear.EndYear,
			},
			Rating: pipe.Rating{
				Score:       v.Node.RatingsSummary.AggregateRating,
				Count:       v.Node.RatingsSummary.VoteCount,
				Certificate: v.Node.Certificate.Rating,
			},
			Duration: v.Node.Runtime.Seconds,
			Genres:   GetRelatedTitleGenres(v.Node.TitleCardGenres.Genres),
		})
	}

	return items
}

func GetRelatedTitleGenres(genres []withText) []pipe.Genre {
	items := []pipe.Genre{}

	for _, v := range genres {
		items = append(items, pipe.Genre{
			Name: v.Text,
			Slug: slug.Make(v.Text),
		})
	}

	return items
}

func GetSoundtracks(edges []soundtrackEdge) []pipe.Soundtrack {
	items := []pipe.Soundtrack{}

	for _, v := range edges {
		items = append(items, pipe.Soundtrack{
			Title:    v.Node.Text,
			Comments: GetSoundtrackComment(v.Node.Comments),
		})
	}

	return items
}

func GetSoundtrackComment(comments []plaidHTMLWrapper) []pipe.SoundtrackComment {
	items := []pipe.SoundtrackComment{}

	for _, v := range comments {
		// Sometimes the plain HTML string equals to useless values
		// we skip those to avoid unnecessary response payloads
		if v.PlaidHTML == "(uncredited)" || v.PlaidHTML == "(Title Theme)" {
			continue
		}

		// Parse plaid HTML string to extract link ID and their name
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(v.PlaidHTML))
		if err != nil {
			items = append(items, pipe.SoundtrackComment{
				HTML: v.PlaidHTML,
			})
			continue
		}

		items = append(items, pipe.SoundtrackComment{
			HTML: v.PlaidHTML,
			Text: doc.Text(),
		})
	}

	return items
}
