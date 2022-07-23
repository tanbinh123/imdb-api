package parser

import (
	"encoding/json"

	"github.com/PuerkitoBio/goquery"
)

// GetNextJSData extracts the second JSON string
// which is usually in the middle or bottom of the web page's source code
// the Next.JS framework automatically generates this.
// Most likely, the structure of this JSON object is going to change.
// But as long as they use the Next.JS framework, the approach remains the same.
func GetNextJSData(doc *goquery.Document) (*nextJSData, error) {
	nextDataJSON := doc.Find("script[type=\"application/json\"][id=\"__NEXT_DATA__\"]").First()
	var nextData nextJSData
	if err := json.Unmarshal([]byte(nextDataJSON.Text()), &nextData); err != nil {
		return nil, err
	}
	return &nextData, nil
}
