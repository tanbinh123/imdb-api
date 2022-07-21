package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTMLToPlainText(input string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		return ""
	}
	return doc.Text()
}
