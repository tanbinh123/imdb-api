package index

type IndexTransform struct {
	ID              string     `json:"id"`
	Validate        validate   `json:"validate"`
	Title           title      `json:"title"`
	Genres          []genre    `json:"genres"`
	Plot            string     `json:"plot"`
	CanHaveEpisodes bool       `json:"canHaveEpisodes"`
	Popularity      popularity `json:"popularity"`
	Images          images     `json:"images"`
	Videos          videos     `json:"videos"`
	Series          series     `json:"series"`
	Reviews         reviews    `json:"reviews"`
	FAQ             faq        `json:"faq"`
	Trivia          trivia     `json:"trivia"`
	Keywords        keyword    `json:"keywords"`
}

type validate struct {
	Type      string `json:"type"`
	IsMovie   bool   `json:"isMovie"`
	IsSeries  bool   `json:"isSeries"`
	IsEpisode bool   `json:"isEpisode"`
	IsAdult   bool   `json:"isAdult"`
}

type title struct {
	Text     string   `json:"text"`
	Original string   `json:"original"`
	AKA      []string `json:"aka"`
}

type genre struct {
	Name string `json:"name"`
	Slug string `json:"slug"` // can directly be used in advances search feature
}

type popularity struct {
	Rank       int64 `json:"rank"`
	Difference int64 `json:"difference"`
}

type videos struct {
	Total int64       `json:"total"`
	Items []videoItem `json:"items"`
}

type videoItem struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int64          `json:"duration"`
	Thumbnail   thumbnail      `json:"thumbnail"`
	Playback    []playbackItem `json:"playback"`
	IsMature    bool           `json:"isMature"`
}

type series struct {
	ID                string        `json:"id"`
	Title             seriesTitle   `json:"title"`
	Current           seriesCurrent `json:"current"`
	NextEpisodeID     string        `json:"nextEpisodeId"`
	PreviousEpisodeID string        `json:"previousEpisodeId"`
	ReleaseYear       releaseYear   `json:"releaseYear"`
}

type seriesTitle struct {
	Text     string `json:"text"`
	Original string `json:"original"`
}

type seriesCurrent struct {
	Episode int64 `json:"episode"` // episode number (starts from 1)
	Season  int64 `json:"season"`  // season number (starts from 1)
}

type releaseYear struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type playbackItem struct {
	Quality  string `json:"quality"`
	MIMEType string `json:"mimeType"`
	URL      string `json:"url"`
}

type thumbnail struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type images struct {
	Total   int64        `json:"total"`
	Primary primaryImage `json:"primary"`
	Items   []imageItem  `json:"items"`
}

type primaryImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type imageItem struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Width   int64  `json:"width"`
	Height  int64  `json:"height"`
	Caption string `json:"caption"`
}

type reviews struct {
	Featured []featuredReviewItem `json:"featured"`
	Users    usersReviews         `json:"users"`
	External externalReviews      `json:"external"`
}

type featuredReviewItem struct {
	ID        string       `json:"id"`
	Author    reviewAuthor `json:"author"`
	Summary   string       `json:"summary"`
	Text      string       `json:"text"`
	Likes     int64        `json:"likes"`
	Dislikes  int64        `json:"dislikes"`
	CreatedAt string       `json:"createdAt"`
}

type reviewAuthor struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Rating   int64  `json:"rating"`
}

type usersReviews struct {
	Total int64 `json:"total"`
}

type externalReviews struct {
	Total int64 `json:"total"`
}

type faq struct {
	Total int64     `json:"total"`
	Items []faqItem `json:"items"`
}

type faqItem struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

type trivia struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}

type keyword struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}
