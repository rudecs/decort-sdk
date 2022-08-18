package decortsdk

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/rudecs/decort-sdk/config"
	"github.com/rudecs/decort-sdk/internal/client"
)

type Client struct {
	decortUrl string
	client    *http.Client
}

func New(cfg config.Config) *Client {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	return &Client{
		decortUrl: cfg.DecortURL,
		client:    client.NewHttpClient(cfg),
	}
}

func (dc *Client) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, method, dc.decortUrl+"/restmachine"+url, body)
	if err != nil {
		return nil, err
	}

	resp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	return respBytes, nil
}
