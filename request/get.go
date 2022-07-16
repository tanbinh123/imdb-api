package request

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Get(url string) (*goquery.Document, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	client := &http.Client{
		// TODO: timeout options
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url(%v): %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to url(%v) failed with status %v %v", url, res.StatusCode, res.Status)
	}

	// TODO: refactor
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url(%v) response: %v", url, err)
	}

	return doc, nil
}
