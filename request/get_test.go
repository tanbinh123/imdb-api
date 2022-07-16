package request

import "testing"

func TestGet(t *testing.T) {
	url := "https://m.imdb.com/title/tt20256528"
	_, err := Get(url)
	if err != nil {
		t.Errorf("Failed to Get URL %v, err: %v.", url, err)
	}
}
