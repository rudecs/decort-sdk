package decortsdk

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/rudecs/decort-sdk/pkg/cloudapi"
	"github.com/rudecs/decort-sdk/pkg/cloudbroker"

	"github.com/google/go-querystring/query"
	"github.com/rudecs/decort-sdk/config"
	"github.com/rudecs/decort-sdk/internal/client"
)

// HTTP-client for platform
type DecortClient struct {
	decortURL string
	client    *http.Client
}

// Сlient builder
func New(cfg config.Config) *DecortClient {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	return &DecortClient{
		decortURL: cfg.DecortURL,
		client:    client.NewHttpClient(cfg),
	}
}

// CloudAPI builder
func (dc *DecortClient) CloudAPI() *cloudapi.CloudAPI {
	return cloudapi.New(dc)
}

// CloudBroker builder
func (dc *DecortClient) CloudBroker() *cloudbroker.CloudBroker {
	return cloudbroker.New(dc)
}

// DecortApiCall method for sending requests to the platform
func (dc *DecortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, method, dc.decortURL+"/restmachine"+url, body)
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
