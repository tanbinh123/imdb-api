package index

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/Scrip7/imdb-api/constants"
	"github.com/Scrip7/imdb-api/request"
)

func Title(id int) (*TitleIndexTransform, error) {
	url := fmt.Sprintf(constants.TITLE_INDEX, id)
	doc, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	script := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var data TitleIndex
	if err := json.Unmarshal([]byte(script.Text()), &data); err != nil {
		return nil, err
	}

	//
	// Begin Transformation
	//
	transform := TitleIndexTransform{
		Images: TitleImages{
			Total: data.Props.PageProps.MainColumnData.TitleMainImages.Total,
			Items: getImageItems(data.Props.PageProps.MainColumnData.TitleMainImages.Edges),
		},
	}

	return &transform, nil
}

func getImageItems(edges []TitleMainImagesEdge) []TitleImageItem {
	var imageItems []TitleImageItem

	for _, v := range edges {
		imageItems = append(imageItems, TitleImageItem{
			ID:      *v.Node.ID,
			Width:   v.Node.Width,
			Height:  v.Node.Height,
			URL:     v.Node.URL,
			Caption: v.Node.Caption.PlainText,
		})
	}

	return imageItems
}
