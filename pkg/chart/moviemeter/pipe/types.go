package pipe

type ChartMovieMeterTransform struct {
	Titles []*ChartMovieMeterItem `json:"titles" extensions:"x-order=001"`
}

type ChartMovieMeterItem struct {
	ID            string                    `json:"id" extensions:"x-order=001"`
	Title         *ChartMovieMeterItemTitle `json:"title" extensions:"x-order=002"`
	Thumbnail     string                    `json:"thumbnail" extensions:"x-order=003"`
	Rating        float64                   `json:"rating" extensions:"x-order=004"`
	RankingChange int64                     `json:"rankingChange" extensions:"x-order=005"` // Could be negative or positive number, determinate chart ranking change
	ReleaseYear   float64                   `json:"releaseYear" extensions:"x-order=006"`
}

type ChartMovieMeterItemTitle struct {
	Text string `json:"text" extensions:"x-order=001"`
	Slug string `json:"slug" extensions:"x-order=002"`
}
