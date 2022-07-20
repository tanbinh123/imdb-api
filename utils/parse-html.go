package utils

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(r io.ReadCloser) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(r)
}
