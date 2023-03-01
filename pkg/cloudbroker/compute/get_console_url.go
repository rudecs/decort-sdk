package compute

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Request struct for get console URL
type GetConsoleURLRequest struct {
	// ID of compute instance to get console for
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq GetConsoleURLRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// GetConsoleURL gets computes console URL
func (c Compute) GetConsoleURL(ctx context.Context, req GetConsoleURLRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/compute/getConsoleUrl"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	result = strings.ReplaceAll(result, "\\", "")

	return result, nil
}
