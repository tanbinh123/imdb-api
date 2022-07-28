package parser

import (
	"encoding/json"
	"fmt"

	"github.com/Scrip7/imdb-api/pkg/title/photos/pipe"
)

func JSONDecode(input *[]byte) (*TitlePhotosAjaxResponse, error) {
	var data TitlePhotosAjaxResponse
	err := json.Unmarshal(*input, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to json decode ajax response: %v", err)
	}

	return &data, nil
}

func GetPhotos(data []Data) []*pipe.TitlePhotoItem {
	items := []*pipe.TitlePhotoItem{}

	for _, v := range data {
		items = append(items, &pipe.TitlePhotoItem{
			ID:        v.Rmconst,
			Width:     v.Width,
			Height:    v.Height,
			Alt:       v.Alt,
			Caption:   v.Caption,
			Thumbnail: v.Src,
		})
	}

	return items
}

func GetMetaData(meta Meta) *pipe.TitlePhotosMeta {
	return &pipe.TitlePhotosMeta{
		CurrentPage:  meta.CurrentPage,
		ItemsPerPage: meta.ItemsPerPage,
		TotalItems:   meta.TotalItems,
		TotalPages:   meta.TotalPages,
		HasNextPage:  meta.ItemsPerPage*meta.CurrentPage < meta.TotalItems,
	}
}
