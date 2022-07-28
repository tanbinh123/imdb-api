package pipe

type TitlePhotosTransform struct {
	Photos []*TitlePhotoItem `json:"photos" extensions:"x-order=001"`
	Meta   *TitlePhotosMeta  `json:"meta" extensions:"x-order=002"`
}

type TitlePhotoItem struct {
	ID        string `json:"id" extensions:"x-order=001"`
	Height    int64  `json:"height" extensions:"x-order=002"`
	Width     int64  `json:"width" extensions:"x-order=003"`
	Alt       string `json:"alt" extensions:"x-order=004"`
	Caption   string `json:"caption" extensions:"x-order=005"`
	Thumbnail string `json:"thumbnail" extensions:"x-order=006"`
}

type TitlePhotosMeta struct {
	CurrentPage  int64 `json:"currentPage"`
	ItemsPerPage int64 `json:"itemsPerPage"`
	TotalItems   int64 `json:"totalItems"`
	TotalPages   int64 `json:"totalPages"`
	HasNextPage  bool  `json:"hasNextPage"`
}
