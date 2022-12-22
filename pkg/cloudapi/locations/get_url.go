package locations

import (
	"context"
	"net/http"
)

// GetURL gets the portal URL
func (l Locations) GetURL(ctx context.Context) (string, error) {
	url := "/cloudapi/locations/getUrl"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
