package client_test

import (
	"encoding/json"
	"testing"

	"github.com/Scrip7/imdb-api/client"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	url := "https://m.imdb.com/title/tt20256528"
	res, err := client.Get(url, &client.GetOptions{})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestParams(t *testing.T) {
	url := "https://httpbin.org/get"
	p := client.Map{
		"first_param_test":  "123456",
		"second_param_test": "654321",
	}
	res, err := client.Get(url, &client.GetOptions{
		Params: &p,
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)

	type HttpBin struct {
		Args map[string]string `json:"args"`
		URL  string            `json:"url"`
	}

	var data HttpBin
	err = json.Unmarshal(*res, &data)
	assert.Nil(t, err, "Failed to Unmarshal HttpBin response")

	assert.Equal(t, data.Args, map[string]string(
		map[string]string{
			"first_param_test":  "123456",
			"second_param_test": "654321",
		},
	))
	assert.Equal(t, data.URL, "https://httpbin.org/get?first_param_test=123456&second_param_test=654321")
}
