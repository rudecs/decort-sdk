package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for revoke user access
type UserRevokeRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Name of the user to remove
	// Required: true
	Username string `url:"userName"`
}

func (crq UserRevokeRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Username == "" {
		return errors.New("validation-error: field UserName can not be empty")
	}

	return nil
}

// UserRevoke revokes user access to the compute
func (c Compute) UserRevoke(ctx context.Context, req UserRevokeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/userRevoke"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
