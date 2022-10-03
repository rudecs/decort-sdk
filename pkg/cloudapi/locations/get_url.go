package locations

import (
	"context"
	"net/http"
)

func (l Locations) GetUrl(ctx context.Context) (string, error) {
	url := "/locations/getUrl"
	prefix := "/cloudapi"

	url = prefix + url
	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
