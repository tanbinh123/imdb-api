package index

type IndexTransform struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Plot       string     `json:"plot"`
	Popularity popularity `json:"popularity"`
	Videos     videos     `json:"videos"`
	Images     images     `json:"images"`
	Reviews    reviews    `json:"reviews"`
	Keywords   keyword    `json:"keywords"`
}

type popularity struct {
	Rank   int64            `json:"rank"`
	Change popularityChange `json:"change"`
}

type popularityChange struct {
	Direction  string `json:"direction"`
	Difference int64  `json:"difference"`
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

type keyword struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}
