package grid

import (
	"context"
	"net/http"
	"strconv"
)

// Status check if current environment is active
func (g Grid) Status(ctx context.Context) (bool, error) {
	url := "/cloudbroker/grid/status"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}

// StatusGET check if current environment is active
func (g Grid) StatusGET(ctx context.Context) (bool, error) {
	url := "/cloudbroker/grid/status"

	res, err := g.client.DecortApiCall(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
