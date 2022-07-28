package parser

type TitlePhotosAjaxResponse struct {
	Data []Data `json:"data"`
	Meta Meta   `json:"meta"`
}

type Data struct {
	Alt     string `json:"alt"`
	Caption string `json:"caption"`
	Height  int64  `json:"height"`
	Href    string `json:"href"`
	Rmconst string `json:"rmconst"`
	Src     string `json:"src"`
	Width   int64  `json:"width"`
}

type Meta struct {
	CurrentPage  int64 `json:"current_page"`
	ItemsPerPage int64 `json:"items_per_page"`
	TotalItems   int64 `json:"total_items"`
	TotalPages   int64 `json:"total_pages"`
}
