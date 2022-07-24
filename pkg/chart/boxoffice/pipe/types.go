package pipe

type BoxOfficeTransform struct {
	Titles []*BoxOfficeItem `json:"titles"`
}

type BoxOfficeItem struct {
	ID                string              `json:"id"`
	Title             *BoxOfficeItemTitle `json:"title"`
	Thumbnail         string              `json:"thumbnail"`
	WeekendGross      string              `json:"weekendGross"`
	TotalGross        string              `json:"totalGross"`
	WeeksSinceRelease string              `json:"weeksSinceRelease"`
}

type BoxOfficeItemTitle struct {
	Text string `json:"text"`
	Slug string `json:"slug"`
}
