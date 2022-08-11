package client

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type decortClient struct {
	decortUrl string
	client    *http.Client
}

func New(config Config) *decortClient {
	return &decortClient{
		decortUrl: config.DecortUrl,
		client:    newHttpClient(config),
	}
}

func (dc *decortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, _ := http.NewRequestWithContext(ctx, method, dc.decortUrl+"/restmachine"+url, body)

	resp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	return respBytes, nil
}
