package index

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/utils"
	"github.com/Scrip7/imdb-api/utils/srcset"
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
	var nextData nextJSData
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

	schemaData, err := getSchemaOrgData(doc)
	if err != nil {
		return nil, err
	}

	nextData, err := getNextJSData(doc)
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
			AKA:      getTitleAKA(main.AKAs.Edges),
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
				ID: above.PrimaryImage.ID,
				// URL:        schemaData.Image,
				URL:        above.PrimaryImage.URL,
				Width:      above.PrimaryImage.Width,
				Height:     above.PrimaryImage.Height,
				Caption:    above.PrimaryImage.Caption.PlainText,
				Thumbnails: getPrimaryImageThumbnails(doc),
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
		Keywords: keywords{
			Total: above.Keywords.Total,
			Items: strings.Split(schemaData.Keywords, ","),
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
		Soundtracks: getSoundtracks(main.Soundtrack.Edges),
		Related:     getRelatedTitles(main.MoreLikeThisTitles.Edges),
	}

	return &transform, nil
}

// getSchemaOrgData extracts the first JSON string, which is usually on top of the web page's source code
// this is manually generated based on "schema.org" schemas,
// and it's unlikely to change in the future.
func getSchemaOrgData(doc *goquery.Document) (*schemaOrgData, error) {
	scriptJSON := doc.Find("script[type=\"application/ld+json\"]").First()
	var data schemaOrgData
	if err := json.Unmarshal([]byte(scriptJSON.Text()), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// getNextJSData extracts the second JSON string
// which is usually in the middle or bottom of the web page's source code
// the Next.JS framework automatically generates this.
// Most likely, the structure of this JSON object is going to change.
// But as long as they use the Next.JS framework, the approach remains the same.
func getNextJSData(doc *goquery.Document) (*nextJSData, error) {
	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData nextJSData
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}
	return &nextData, nil
}

func getTitleAKA(edges []akaEdge) []string {
	items := []string{}

	for _, v := range edges {
		items = append(items, v.Node.Text)
	}

	return items
}

func getGenres(genres []withTextAndID) []genre {
	items := []genre{}

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

func getPrimaryImageThumbnails(doc *goquery.Document) []primaryImageThumbnail {
	items := []primaryImageThumbnail{}

	poster := doc.Find("section.ipc-page-section div[data-testid=\"hero-media__poster\"] img.ipc-image").First()
	srcsetRaw, exists := poster.Attr("srcset")
	if !exists {
		return items
	}
	srcsetSource := srcset.Parse(srcsetRaw)
	for _, v := range srcsetSource {
		items = append(items, primaryImageThumbnail{
			URL:   v.URL,
			Width: *v.Width,
			// IMDb does not return the "height" or "density"
		})
	}

	return items
}

func getImageItems(edges []imageEdge) []imageItem {
	items := []imageItem{}

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
	items := []videoItem{}

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
	items := []playbackItem{}

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
	items := []featuredReviewItem{}

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
	items := []faqItem{}

	for _, v := range edges {
		items = append(items, faqItem{
			ID:       v.Node.ID,
			Question: v.Node.Question.PlainText,
		})
	}

	return items
}

func getTriviaItems(edges []triviaEdge) []string {
	items := []string{}

	for _, v := range edges {
		items = append(items, utils.ParseHTMLToString(v.Node.Text.PlaidHTML))
	}

	return items
}

func getRelatedTitles(edges []moreLikeThisTitlesEdge) []relatedTitle {
	items := []relatedTitle{}

	for _, v := range edges {
		items = append(items, relatedTitle{
			ID: v.Node.ID,
			Title: relatedTitleName{
				Text:     v.Node.TitleText.Text,
				Original: v.Node.OriginalTitleText.Text,
				Slug:     slug.Make(v.Node.OriginalTitleText.Text),
			},
			Type:            strcase.ToLowerCamel(v.Node.TitleType.ID),
			CanHaveEpisodes: v.Node.CanHaveEpisodes,
			Poster: relatedTitlePoster{
				ID:     v.Node.PrimaryImage.ID,
				Width:  v.Node.PrimaryImage.Width,
				Height: v.Node.PrimaryImage.Height,
				URL:    v.Node.PrimaryImage.URL,
			},
			ReleaseYear: releaseYear{
				From: v.Node.ReleaseYear.Year,
				To:   v.Node.ReleaseYear.EndYear,
			},
			Rating: rating{
				Score:       v.Node.RatingsSummary.AggregateRating,
				Count:       v.Node.RatingsSummary.VoteCount,
				Certificate: v.Node.Certificate.Rating,
			},
			Duration: v.Node.Runtime.Seconds,
			Genres:   getRelatedTitleGenres(v.Node.TitleCardGenres.Genres),
		})
	}

	return items
}

func getRelatedTitleGenres(genres []withText) []genre {
	items := []genre{}

	for _, v := range genres {
		items = append(items, genre{
			Name: v.Text,
			Slug: slug.Make(v.Text),
		})
	}

	return items
}

func getSoundtracks(edges []soundtrackEdge) []soundtrack {
	items := []soundtrack{}

	for _, v := range edges {
		items = append(items, soundtrack{
			Title:    v.Node.Text,
			Comments: getSoundtrackComment(v.Node.Comments),
		})
	}

	return items
}

func getSoundtrackComment(comments []plaidHTMLWrapper) []soundtrackComment {
	items := []soundtrackComment{}

	pushItem := func(item soundtrackComment) {
		items = append(items, item)
	}

	for _, v := range comments {
		// Sometimes the plain HTML string equals to "(uncredited)" for unknown reason
		// Here we double-check to avoid future errors
		// Even though it won't interfere with our business logic below
		if v.PlaidHTML == "(uncredited)" {
			continue
		}

		// Parse plaid HTML string to extract link ID and their name
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(v.PlaidHTML))
		if err != nil {
			continue
		}

		link := doc.Find("a").First()
		href, exists := link.Attr("href")
		if !exists {
			pushItem(soundtrackComment{
				Original: doc.Text(),
			})
			continue
		}

		linkName := strings.TrimSpace(link.Text())
		// The plaid HTML text is usually "Performed by [LINK_NAME]"
		// Here, we replace the link's name with an empty string to extract the headline from it
		// Which would be "Performed by "
		headline := strings.TrimSpace(strings.Replace(doc.Text(), linkName, "", 1))

		// The "href" variable contains a string that contains the ID we need
		// We use the regex below to extract that ID
		// Example: "/name/nm0123456/?ref_=tt_trv_snd"
		r, err := regexp.Compile("(nm[0-9]+)")
		if err != nil {
			pushItem(soundtrackComment{
				Original: doc.Text(),
			})
			log.Err(err).Msg("failed to compile the extract link ID regex")
			continue
		}

		items = append(items, soundtrackComment{
			Original: doc.Text(),
			Parsed: soundTrackCommentParsed{
				Headline: headline,
				Link: linkWrapper{
					ID:   r.FindString(href),
					Name: linkName,
				},
			},
		})
	}

	return items
}
