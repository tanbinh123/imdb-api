package pipe

type IndexTransform struct {
	ID          string         `json:"id"`
	Validate    Validate       `json:"validate"`
	Title       Title          `json:"title"`
	Genres      []Genre        `json:"genres"`
	Plot        string         `json:"plot"`
	Popularity  Popularity     `json:"popularity"`
	Images      Images         `json:"images"`
	Videos      Videos         `json:"videos"`
	Reviews     Reviews        `json:"reviews"`
	FAQ         FAQ            `json:"faq"`
	Trivia      Trivia         `json:"trivia"`
	Keywords    Keywords       `json:"keywords"`
	Series      Series         `json:"series"` // only when viewing an episode of a series
	Soundtracks []Soundtrack   `json:"soundtracks"`
	Related     []RelatedTitle `json:"related"` // list of related titles
}

type Validate struct {
	Type            string `json:"type"`
	IsMovie         bool   `json:"isMovie"`
	IsSeries        bool   `json:"isSeries"`
	IsEpisode       bool   `json:"isEpisode"`
	IsAdult         bool   `json:"isAdult"`
	CanHaveEpisodes bool   `json:"canHaveEpisodes"`
}

type Title struct {
	Text     string   `json:"text"`
	Original string   `json:"original"`
	Slug     string   `json:"slug"`
	AKA      []string `json:"aka"`
}

type Genre struct {
	Name string `json:"name"`
	Slug string `json:"slug"` // can directly be used in advances search feature
}

type Popularity struct {
	Rank       int64 `json:"rank"`
	Difference int64 `json:"difference"`
}

type Videos struct {
	Total     int64              `json:"total"`
	Primaries []VideoItemPrimary `json:"primaries"`
	Items     []VideoItem        `json:"items"`
}

type VideoItemPrimary struct {
	ID          string           `json:"id"`
	Type        VideoTypeWrapper `json:"type"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Duration    int64            `json:"duration"`
	Thumbnail   Thumbnail        `json:"thumbnail"`
	Playback    []PlaybackItem   `json:"playback"`
	IsMature    bool             `json:"isMature"`
}

type VideoItem struct {
	ID        string           `json:"id"`
	Type      VideoTypeWrapper `json:"type"`
	Title     string           `json:"title"`
	Duration  int64            `json:"duration"`
	Thumbnail Thumbnail        `json:"thumbnail"`
}

type VideoTypeWrapper struct {
	Text string `json:"text"`
	Slug string `json:"slug"`
}

type Series struct {
	ID          SeriesID      `json:"id"`
	Title       SeriesTitle   `json:"title"`
	Current     SeriesCurrent `json:"current"`
	ReleaseYear ReleaseYear   `json:"releaseYear"`
}

type SeriesID struct {
	Parent          string `json:"parent"`
	EpisodeNext     string `json:"episodeNext"`
	EpisodePrevious string `json:"episodePrevious"`
}

type SeriesTitle struct {
	Text     string `json:"text"`
	Original string `json:"original"`
	Slug     string `json:"slug"`
}

type SeriesCurrent struct {
	Episode int64 `json:"episode"` // episode number (starts from 1)
	Season  int64 `json:"season"`  // season number (starts from 1)
}

type ReleaseYear struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type PlaybackItem struct {
	Quality  string `json:"quality"`
	MIMEType string `json:"mimeType"`
	URL      string `json:"url"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Images struct {
	Total   int64        `json:"total"`
	Primary PrimaryImage `json:"primary"`
	Items   []ImageItem  `json:"items"`
}

type PrimaryImage struct {
	ID         string                  `json:"id"`
	URL        string                  `json:"url"`
	Width      int64                   `json:"width"`
	Height     int64                   `json:"height"`
	Caption    string                  `json:"caption"`
	Thumbnails []PrimaryImageThumbnail `json:"thumbnails"`
}

type PrimaryImageThumbnail struct {
	URL   string `json:"url"`
	Width int64  `json:"width"`
}

type ImageItem struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Width   int64  `json:"width"`
	Height  int64  `json:"height"`
	Caption string `json:"caption"`
}

type Reviews struct {
	Featured []FeaturedReviewItem `json:"featured"`
	Users    UsersReviews         `json:"users"`
	External ExternalReviews      `json:"external"`
}

type FeaturedReviewItem struct {
	ID        string       `json:"id"`
	Author    ReviewAuthor `json:"author"`
	Summary   string       `json:"summary"`
	Text      string       `json:"text"`
	Likes     int64        `json:"likes"`
	Dislikes  int64        `json:"dislikes"`
	CreatedAt string       `json:"createdAt"`
}

type ReviewAuthor struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Rating   int64  `json:"rating"`
}

type UsersReviews struct {
	Total int64 `json:"total"`
}

type ExternalReviews struct {
	Total int64 `json:"total"`
}

type FAQ struct {
	Total int64     `json:"total"` // total number of keywords that are related to this title
	Items []FAQItem `json:"items"` // a short list of FAQs as preview
}

type FAQItem struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

type Trivia struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}

type Keywords struct {
	Total int64    `json:"total"` // total number of keywords that are related to this title
	Items []string `json:"items"` // a short list of keywords as preview
}

type Soundtrack struct {
	Title    string              `json:"title"`
	Comments []SoundtrackComment `json:"comments"`
}

type SoundtrackComment struct {
	HTML string `json:"html"`
	Text string `json:"text"`
}

type RelatedTitle struct {
	ID              string             `json:"id"`
	Title           RelatedTitleName   `json:"title"`
	Type            string             `json:"type"`
	CanHaveEpisodes bool               `json:"canHaveEpisodes"`
	Poster          RelatedTitlePoster `json:"poster"`
	ReleaseYear     ReleaseYear        `json:"releaseYear"`
	Rating          Rating             `json:"rating"`
	Duration        int64              `json:"duration"` // unit is seconds
	Genres          []Genre            `json:"genres"`
}

type RelatedTitleName struct {
	Text     string `json:"text"`
	Original string `json:"original"`
	Slug     string `json:"slug"`
}

type RelatedTitlePoster struct {
	ID     string `json:"id"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
}

type Rating struct {
	Score       float64 `json:"score"`       // IMDb score from 0 to 10 (includes precision)
	Count       int64   `json:"count"`       // Voters count
	Certificate string  `json:"certificate"` // For example "TV-14"
}
