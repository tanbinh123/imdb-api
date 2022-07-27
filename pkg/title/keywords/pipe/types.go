package pipe

type TitleKeywordsTransform struct {
	Keywords []*TitleKeyword `json:"keywords"`
	// TODO:
}

type TitleKeyword struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Slug string `json:"slug"`
}
