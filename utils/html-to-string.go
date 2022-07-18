package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTMLToString(input string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		return "__FAILED_TO_PARSE_HTML_TEXT__"
	}
	return doc.Text()
}
