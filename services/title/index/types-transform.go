package index

type TitleIndexTransform struct {
	Images TitleImages `json:"images"`
}

type TitleImages struct {
	Total int64            `json:"total"`
	Items []TitleImageItem `json:"items"`
}

type TitleImageItem struct {
	ID      string `json:"id"`
	Width   int64  `json:"width"`
	Height  int64  `json:"height"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}
