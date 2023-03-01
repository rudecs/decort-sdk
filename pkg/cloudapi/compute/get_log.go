package compute

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for get compute logs
type GetLogRequest struct {
	// ID of compute instance to get log for
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Path to log file
	// Required: true
	Path string `url:"path" json:"path"`
}

func (crq GetLogRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Path == "" {
		return errors.New("validation-error: field Path can not be empty")
	}

	return nil
}

// GetLog gets compute's log file by path
func (c Compute) GetLog(ctx context.Context, req GetLogRequest) (string, error) {
	err := req.validate()
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

// GetLogGet gets compute's log file by path
func (c Compute) GetLogGet(ctx context.Context, req GetLogRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi//compute/getLog"

	res, err := c.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
