package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/services/title/index/pipe"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
)

func GetTitle(text string, original string, akaEdges []akaEdge) *pipe.Title {
	return &pipe.Title{
		Text:     text,
		Original: original,
		Slug:     slug.Make(original),
		AKA:      getTitleAKA(akaEdges),
	}
}

func GetGenres(genres []withTextAndID) []*pipe.Genre {
	items := []*pipe.Genre{}

	for _, v := range genres {
		items = append(items, &pipe.Genre{
			Name: v.Text,
			Slug: slug.Make(v.ID),
		})
	}

	return items
}

func GetPopularity(m meterRanking) *pipe.Popularity {
	return &pipe.Popularity{
		Rank:       m.CurrentRank,
		Difference: getRankingDifference(m.RankChange),
	}
}

func GetImages(total int64, primaryImage primaryImage, edges []imageEdge, doc *goquery.Document) *pipe.Images {
	return &pipe.Images{
		Total: total,
		Primary: &pipe.PrimaryImage{
			ID:         primaryImage.ID,
			URL:        primaryImage.URL,
			Width:      primaryImage.Width,
			Height:     primaryImage.Height,
			Caption:    primaryImage.Caption.PlainText,
			Thumbnails: getPrimaryImageThumbnails(doc),
		},
		Items: getImageItems(edges),
	}
}

func GetVideos(total int64, primaryEdges []primaryVideosEdge, stripEdges []videoStripEdge) *pipe.Videos {
	return &pipe.Videos{
		Total:     total,
		Primaries: getPrimaryVideos(primaryEdges),
		Items:     getVideoItems(stripEdges),
	}
}

func GetCast(edges []castNodeWrapper) []*pipe.Cast {
	items := []*pipe.Cast{}

	for _, v := range edges {
		items = append(items, &pipe.Cast{
			ID:   v.Node.Name.ID,
			Name: v.Node.Name.NameText.Text,
			Slug: slug.Make(v.Node.Name.NameText.Text),
			Image: &pipe.CastImage{
				URL:    v.Node.Name.PrimaryImage.URL,
				Width:  v.Node.Name.PrimaryImage.Width,
				Height: v.Node.Name.PrimaryImage.Width,
			},
			Characters: getCastChapters(v.Node.Characters),
			EpisodeCredits: &pipe.CastEpisodeCredits{
				Total: v.Node.EpisodeCredits.Total,
				Years: pipe.YearRange{
					From: v.Node.EpisodeCredits.YearRange.Year,
					To:   v.Node.EpisodeCredits.YearRange.EndYear,
				},
			},
		})
	}

	return items
}

func GetReviews(totalUsers int64, totalExternal int64, featuresEdges []featuredReviewEdge) *pipe.Reviews {
	return &pipe.Reviews{
		Featured: getFeaturedReviews(featuresEdges),
		Users: &pipe.UsersReviews{
			Total: totalUsers,
		},
		External: &pipe.ExternalReviews{
			Total: totalExternal,
		},
	}
}

func GetFAQ(total int64, edges []faqEdge) *pipe.FAQ {
	return &pipe.FAQ{
		Total: total,
		Items: getFAQItems(edges),
	}
}

func GetTrivia(total int64, edges []triviaEdge) *pipe.Trivia {
	return &pipe.Trivia{
		Total: total,
		Items: getTriviaItems(edges),
	}
}

func GetKeywords(k keywords) *pipe.Keywords {
	return &pipe.Keywords{
		Total: k.Total,
		Items: getKeywordItems(k.Edges),
	}
}

func GetSeries(s series) *pipe.Series {
	return &pipe.Series{
		ID: &pipe.SeriesID{
			Parent:          s.Series.ID,
			EpisodeNext:     s.NextEpisode.ID,
			EpisodePrevious: s.PreviousEpisode.ID,
		},
		Title: &pipe.SeriesTitle{
			Text:     s.Series.TitleText.Text,
			Original: s.Series.OriginalTitleText.Text,
			Slug:     slug.Make(s.Series.OriginalTitleText.Text),
		},
		Current: &pipe.SeriesCurrent{
			Episode: s.EpisodeNumber.EpisodeNumber,
			Season:  s.EpisodeNumber.SeasonNumber,
		},
		ReleaseYear: &pipe.YearRange{
			From: s.Series.ReleaseYear.Year,
			To:   s.Series.ReleaseYear.EndYear,
		},
	}
}

func GetSoundtracks(edges []soundtrackEdge) []*pipe.Soundtrack {
	items := []*pipe.Soundtrack{}

	for _, v := range edges {
		items = append(items, &pipe.Soundtrack{
			Title:    v.Node.Text,
			Comments: getSoundtrackComment(v.Node.Comments),
		})
	}

	return items
}

func GetRelatedTitles(edges []moreLikeThisTitlesEdge) []*pipe.RelatedTitle {
	items := []*pipe.RelatedTitle{}

	for _, v := range edges {
		items = append(items, &pipe.RelatedTitle{
			ID: v.Node.ID,
			Title: &pipe.RelatedTitleName{
				Text:     v.Node.TitleText.Text,
				Original: v.Node.OriginalTitleText.Text,
				Slug:     slug.Make(v.Node.OriginalTitleText.Text),
			},
			Type:            strcase.ToLowerCamel(v.Node.TitleType.ID),
			CanHaveEpisodes: v.Node.CanHaveEpisodes,
			Poster: &pipe.RelatedTitlePoster{
				ID:     v.Node.PrimaryImage.ID,
				Width:  v.Node.PrimaryImage.Width,
				Height: v.Node.PrimaryImage.Height,
				URL:    v.Node.PrimaryImage.URL,
			},
			ReleaseYear: &pipe.YearRange{
				From: v.Node.ReleaseYear.Year,
				To:   v.Node.ReleaseYear.EndYear,
			},
			Rating: &pipe.Rating{
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
