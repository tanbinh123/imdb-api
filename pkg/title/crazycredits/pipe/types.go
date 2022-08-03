package pipe

type TitleCrazyCreditsTransform struct {
	Items []*CrazyCreditItem `json:"items"`
}

type CrazyCreditItem struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
