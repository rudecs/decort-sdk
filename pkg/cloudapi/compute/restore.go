package compute

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for restore compute
type RestoreRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq RestoreRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Restore restore compute from recycle bin
func (c Compute) Restore(ctx context.Context, req RestoreRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/compute/restore"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
