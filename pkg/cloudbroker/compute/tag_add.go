package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add tag to compute
type TagAddRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds"`

	// Tag key
	// Required: true
	Key string `url:"key" json:"key"`

	// Tag value
	// Required: true
	Value string `url:"value" json:"value"`
}

func (crq TagAddRequest) validate() error {
	if len(crq.ComputeIDs) == 0 {
		return errors.New("validation-error: field ComputeIDs must be set")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key must be set")
	}
	if crq.Value == "" {
		return errors.New("validation-error: field Value must be set")
	}

	return nil
}

// TagAdd add tag to compute tags dict
func (c Compute) TagAdd(ctx context.Context, req TagAddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/tagAdd"

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
