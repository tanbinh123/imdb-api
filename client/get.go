package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/v7/httpclient"
)

type linearBackoff struct {
	backoffInterval int
}

func (lb *linearBackoff) Next(retry int) time.Duration {
	if retry <= 0 {
		return 0 * time.Millisecond
	}
	return time.Duration(retry*lb.backoffInterval) * time.Millisecond
}

var (
	client = httpclient.NewClient(
		httpclient.WithHTTPTimeout(4*time.Second),
		httpclient.WithRetrier(heimdall.NewRetrier(&linearBackoff{100})),
		httpclient.WithRetryCount(4),
	)
)

func Get(url string) (*io.ReadCloser, error) {
	res, err := client.Get(url, http.Header{
		// TODO: dynamic language
		"accept-language": {"en-US,en;q=0.9"},
		// TODO: dynamic user-agent
		"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url(%v): %v", url, err)
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to url(%v) failed with status %v", url, res.Status)
	}

	return &res.Body, nil
}
