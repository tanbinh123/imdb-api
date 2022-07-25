package pipe

type ChartCommonTransform struct {
	Type   string      `json:"type" extensions:"x-order=001"`
	Titles []*ChatItem `json:"titles" extensions:"x-order=002"`
}

type ChatItem struct {
	ID          string          `json:"id" extensions:"x-order=001"`
	Title       *ChartItemTitle `json:"title" extensions:"x-order=002"`
	Thumbnail   string          `json:"thumbnail" extensions:"x-order=003"`
	Rating      float64         `json:"rating" extensions:"x-order=004"`
	ReleaseYear float64         `json:"releaseYear" extensions:"x-order=005"`
}

type ChartItemTitle struct {
	Text string `json:"text"`
	Slug string `json:"slug"`
}
