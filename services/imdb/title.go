package imdb

import (
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Scrip7/imdb-api/request"
	"github.com/Scrip7/imdb-api/utils"
	"github.com/gosimple/slug"
)

type TitleRating struct {
	Count int `json:"count"`
	// Best    int     `json:"best"`
	// Worst   int     `json:"worst"`
	Rating  float64 `json:"rating"`
	Content string  `json:"content"`
}

type TitleGenre struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleActor struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// URL  string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleDirector struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// URL  string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleCreator struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// URL  string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleRelatedItem struct {
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Slug   string  `json:"slug"`
	Name   string  `json:"name"`
	Poster string  `json:"poster"`
	Rating float64 `json:"rating"`
}

type TitleCountry struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleLanguage struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TitleTransform struct {
	IsSeries      bool               `json:"isSeries"`
	ID            string             `json:"id"`
	URL           string             `json:"url"`
	Slug          string             `json:"slug"`
	Name          string             `json:"name"`
	AlternateName string             `json:"nameAlternate"`
	Poster        string             `json:"poster"`
	Description   string             `json:"description"`
	Duration      float64            `json:"duration"`
	Rating        TitleRating        `json:"rating"`
	PublishDate   string             `json:"publishDate"`
	Genres        []TitleGenre       `json:"genres"`
	Keywords      []string           `json:"keywords"`
	Languages     []TitleLanguage    `json:"languages"`
	Countries     []TitleCountry     `json:"countries"`
	Related       []TitleRelatedItem `json:"relatedTitles"`
}

type TitleResponse struct {
	Context       string `json:"@context"`
	Type          string `json:"@type"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	AlternateName string `json:"alternateName"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	Review        struct {
		Type         string `json:"@type"`
		ItemReviewed struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"itemReviewed"`
		Author struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"author"`
		DateCreated  string `json:"dateCreated"`
		InLanguage   string `json:"inLanguage"`
		Name         string `json:"name"`
		ReviewBody   string `json:"reviewBody"`
		ReviewRating struct {
			Type        string `json:"@type"`
			WorstRating int    `json:"worstRating"`
			BestRating  int    `json:"bestRating"`
			RatingValue int    `json:"ratingValue"`
		} `json:"reviewRating"`
	} `json:"review"`
	AggregateRating struct {
		Type        string  `json:"@type"`
		RatingCount int     `json:"ratingCount"`
		BestRating  int     `json:"bestRating"`
		WorstRating int     `json:"worstRating"`
		RatingValue float64 `json:"ratingValue"`
	} `json:"aggregateRating"`
	ContentRating string   `json:"contentRating"`
	Genre         []string `json:"genre"`
	DatePublished string   `json:"datePublished"`
	Keywords      string   `json:"keywords"`
	Trailer       struct {
		Type      string `json:"@type"`
		Name      string `json:"name"`
		EmbedURL  string `json:"embedUrl"`
		Thumbnail struct {
			Type       string `json:"@type"`
			ContentURL string `json:"contentUrl"`
		} `json:"thumbnail"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Description  string `json:"description"`
	} `json:"trailer"`
	Duration string `json:"duration"`
}

func Title(id string) (*TitleTransform, error) {
	url := fmt.Sprintf("https://m.imdb.com/title/tt%v", utils.ExtractNumbers(id))
	doc, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	script := doc.Find("script[type=\"application/ld+json\"]").First()
	var data TitleResponse
	if err := json.Unmarshal([]byte(script.Text()), &data); err != nil {
		return nil, err
	}

	// Start transforming
	transform := TitleTransform{
		IsSeries:      strings.ToLower(data.Type) == "tvseries",
		ID:            utils.ExtractNumbers(data.URL),
		URL:           data.URL,
		Slug:          slug.Make(data.Name),
		Name:          data.Name,
		AlternateName: data.AlternateName,
		Poster:        data.Image,
		Description:   strings.TrimSpace(html.UnescapeString(data.Description)),
		Rating: TitleRating{
			Count:   data.AggregateRating.RatingCount,
			Rating:  data.AggregateRating.RatingValue,
			Content: data.ContentRating,
			// Best: data.AggregateRating.BestRating,
			// Worst: data.AggregateRating.WorstRating,
		},
		Keywords:    strings.Split(data.Keywords, ","),
		PublishDate: data.DatePublished,
	}

	if data.Duration != "" {
		h, err := utils.ParseDuration(data.Duration)
		if err != nil {
			return nil, fmt.Errorf("failed to parse duration: %v", err)
		}
		transform.Duration = h.Seconds()
	}

	for _, v := range data.Genre {
		transform.Genres = append(transform.Genres, TitleGenre{
			Name: v,
			Slug: slug.Make(v),
		})
	}

	// Languages
	languageLinks := doc.Find("li[data-testid=\"title-details-languages\"] ul li a")
	languageLinks.Each(func(i int, s *goquery.Selection) {
		transform.Languages = append(transform.Languages, TitleLanguage{
			Name: s.Text(),
			Slug: slug.Make(s.Text()),
		})
	})

	// Countries
	doc.Find("li[data-testid=\"title-details-origin\"] ul li a").Each(func(i int, s *goquery.Selection) {
		transform.Countries = append(transform.Countries, TitleCountry{
			Name: s.Text(),
			Slug: slug.Make(s.Text()),
		})
	})

	// Related posts (more like this)
	relatedSection := doc.Find("section[data-testid=\"MoreLikeThis\"]")
	relatedItems := relatedSection.Find("div.ipc-poster-card")
	relatedItems.Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a.ipc-poster-card__title").Attr("href")
		url := utils.RemoveQueryString(link)

		img := s.Find("img.ipc-image")
		name, _ := img.Attr("alt")
		poster, _ := img.Attr("src")
		rating := strings.TrimSpace(s.Find("span.ipc-rating-star").Text())
		rate, _ := strconv.ParseFloat(rating, 64)

		transform.Related = append(transform.Related, TitleRelatedItem{
			ID:     utils.ExtractNumbers(url),
			URL:    url,
			Name:   name,
			Slug:   slug.Make(name),
			Poster: poster,
			Rating: rate,
		})
	})

	// TODO: trailers

	return &transform, nil
}