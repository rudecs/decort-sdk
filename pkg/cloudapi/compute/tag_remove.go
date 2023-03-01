package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove tag from compute
type TagRemoveRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Tag key
	// Required: true
	Key string `url:"key" json:"key"`
}

func (crq TagRemoveRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}

	return nil
}

// TagRemove removes tag from compute tags dict
func (c Compute) TagRemove(ctx context.Context, req TagRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/tagRemove"

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
