package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gojek/heimdall/v7"
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

var client = httpclient.NewClient(
	httpclient.WithHTTPTimeout(4*time.Second),
	httpclient.WithRetrier(heimdall.NewRetrier(&linearBackoff{
		backoffInterval: 100,
	})),
	httpclient.WithRetryCount(4),
)

type GetOptions struct {
	// Language string
	Params *Map
}

type Map map[string]string

func Get(url string, options *GetOptions) (*[]byte, error) {
	if options.Params != nil && len(*options.Params) > 0 {
		base, err := u.Parse(url)
		if err != nil {
			return nil, err
		}

		// Query params
		params := u.Values{}
		for k, v := range *options.Params {
			params.Add(k, v)
		}
		base.RawQuery = params.Encode()

		url = base.String()
	}

	res, err := client.Get(url, http.Header{
		// TODO: dynamic language
		"accept-language": {"en-US,en;q=0.9"},
		// TODO: dynamic user-agent
		"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url(%v): %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to url(%v) failed with status %v", url, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return &body, nil
}
