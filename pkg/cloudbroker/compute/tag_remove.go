package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove tag from compute
type TagRemoveRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds"`

	// Tag key
	// Required: true
	Key string `url:"key" json:"key"`
}

func (crq TagRemoveRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key must be set")
	}

	return nil
}

// TagRemove removes tag from compute tags dict
func (c Compute) TagRemove(ctx context.Context, req TagRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/tagRemove"

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
