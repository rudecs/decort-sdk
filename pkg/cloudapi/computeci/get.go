package computeci

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for information about computeci
type GetRequest struct {
	// ID of the Compute CI
	// Required: true
	ComputeCIID uint64 `url:"computeciId" json:"computeciId"`
}

func (krq GetRequest) validate() error {
	if krq.ComputeCIID == 0 {
		return errors.New("field ComputeCIID can not be empty or equal to 0")
	}

	return nil
}

// Get gets information about computeci by ID
func (c ComputeCI) Get(ctx context.Context, req GetRequest) (*ItemComputeCI, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/computeci/get"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ItemComputeCI{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
