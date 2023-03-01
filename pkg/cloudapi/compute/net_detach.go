package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for detach networ to compute
type NetDetachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// IP of the network interface
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// MAC of the network interface
	// Required: false
	MAC string `url:"mac,omitempty" json:"mac,omitempty"`
}

func (crq NetDetachRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// NetDetach detach network to compute
func (c Compute) NetDetach(ctx context.Context, req NetDetachRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/netDetach"

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
