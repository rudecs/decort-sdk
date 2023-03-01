package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add tag to compute
type TagAddRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Tag key
	// Required: true
	Key string `url:"key" json:"key"`

	// Tag value
	// Required: true
	Value string `url:"value" json:"value"`
}

func (crq TagAddRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}
	if crq.Value == "" {
		return errors.New("validation-error: field Value can not be empty")
	}

	return nil
}

// TagAdd add tag to compute tags dict
func (c Compute) TagAdd(ctx context.Context, req TagAddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/tagAdd"

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
