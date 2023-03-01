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
	"github.com/rudecs/decort-sdk/pkg/cloudapi"
	"github.com/rudecs/decort-sdk/pkg/cloudbroker"
)

// Legacy HTTP-client for platform
type LegacyDecortClient struct {
	decortURL string
	client    *http.Client
}

// Legacy client builder
func NewLegacy(cfg config.LegacyConfig) *LegacyDecortClient {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	return &LegacyDecortClient{
		decortURL: cfg.DecortURL,
		client:    client.NewLegacyHttpClient(cfg),
	}
}

// CloudAPI builder
func (ldc *LegacyDecortClient) CloudAPI() *cloudapi.CloudAPI {
	return cloudapi.New(ldc)
}

// CloudBroker builder
func (ldc *LegacyDecortClient) CloudBroker() *cloudbroker.CloudBroker {
	return cloudbroker.New(ldc)
}

// DecortApiCall method for sending requests to the platform
func (ldc *LegacyDecortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, method, ldc.decortURL+"/restmachine"+url, body)
	if err != nil {
		return nil, err
	}

	resp, err := ldc.client.Do(req)
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
