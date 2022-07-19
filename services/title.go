package services

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"github.com/gosimple/slug"

	"github.com/Scrip7/imdb-api/client"
	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/utils"
)

type titleRating struct {
	Count int `json:"count"`
	// Best    int     `json:"best"`
	// Worst   int     `json:"worst"`
	Rating  float64 `json:"rating"`
	Content string  `json:"content"`
}

type titleGenre struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type titleRelatedItem struct {
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Slug   string  `json:"slug"`
	Name   string  `json:"name"`
	Poster string  `json:"poster"`
	Rating float64 `json:"rating"`
}

type titleCountry struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type titleLanguage struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type titleTransform struct {
	IsSeries      bool               `json:"isSeries"`
	ID            string             `json:"id"`
	URL           string             `json:"url"`
	Slug          string             `json:"slug"`
	Name          string             `json:"name"`
	AlternateName string             `json:"nameAlternate"`
	Poster        string             `json:"poster"`
	Description   string             `json:"description"`
	Duration      float64            `json:"duration"`
	Rating        titleRating        `json:"rating"`
	PublishDate   string             `json:"publishDate"`
	Genres        []titleGenre       `json:"genres"`
	Keywords      []string           `json:"keywords"`
	Languages     []titleLanguage    `json:"languages"`
	Countries     []titleCountry     `json:"countries"`
	Related       []titleRelatedItem `json:"relatedTitles"`
}

type titleResponse struct {
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

func Title(id int) (*titleTransform, error) {
	url := fmt.Sprintf(constants.URL_TITLE_INDEX, id)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := utils.ParseHTML(*res)
	if err != nil {
		return nil, err
	}

	script := doc.Find("script[type=\"application/ld+json\"]").First()
	var data titleResponse
	if err := json.Unmarshal([]byte(script.Text()), &data); err != nil {
		return nil, err
	}

	// Start transforming
	transform := titleTransform{
		IsSeries:      strings.ToLower(data.Type) == "tvseries",
		ID:            utils.ExtractNumbers(data.URL),
		URL:           data.URL,
		Slug:          slug.Make(data.Name),
		Name:          data.Name,
		AlternateName: data.AlternateName,
		Poster:        data.Image,
		Description:   strings.TrimSpace(html.UnescapeString(data.Description)),
		Rating: titleRating{
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
		transform.Genres = append(transform.Genres, titleGenre{
			Name: v,
			Slug: slug.Make(v),
		})
	}

	// Languages
	languageLinks := doc.Find("li[data-testid=\"title-details-languages\"] ul li a")
	languageLinks.Each(func(i int, s *goquery.Selection) {
		transform.Languages = append(transform.Languages, titleLanguage{
			Name: s.Text(),
			Slug: slug.Make(s.Text()),
		})
	})

	// Countries
	doc.Find("li[data-testid=\"title-details-origin\"] ul li a").Each(func(i int, s *goquery.Selection) {
		transform.Countries = append(transform.Countries, titleCountry{
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

		transform.Related = append(transform.Related, titleRelatedItem{
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
