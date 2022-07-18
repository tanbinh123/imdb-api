package index

type IndexTransform struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Plot       string     `json:"plot"`
	Popularity popularity `json:"popularity"`
	Videos     videos     `json:"videos"`
	Images     images     `json:"images"`
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

type keyword struct {
	Total int64    `json:"total"`
	Items []string `json:"items"`
}
