package compute

import (
	"context"
	"errors"
	"net/http"
)

type GetLogRequest struct {
	ComputeID uint64 `url:"computeId"`
	Path      string `url:"path"`
}

func (crq GetLogRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.Path == "" {
		return errors.New("validation-error: field Path can not be empty")
	}

	return nil
}

func (c Compute) GetLog(ctx context.Context, req GetLogRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/compute/getLog"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c Compute) GetLogGet(ctx context.Context, req GetLogRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/compute/getLog"
	prefix := "/cloudapi"

	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
