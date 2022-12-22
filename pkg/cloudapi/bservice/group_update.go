package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update existing Compute group
type GroupUpdateRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId"`

	// Specify non-empty string to update Compute Group name
	// Required: false
	Name string `url:"name,omitempty"`

	// Specify non-empty string to update group role
	// Required: false
	Role string `url:"role,omitempty"`

	// Specify positive value to set new compute CPU count
	// Required: false
	CPU uint64 `url:"cpu,omitempty"`

	// Specify positive value to set new compute RAM volume in MB
	// Required: false
	RAM uint64 `url:"ram,omitempty"`

	// Specify new compute boot disk size in GB
	// Required: false
	Disk uint64 `url:"disk,omitempty"`

	// Force resize Compute Group
	// Required: false
	Force bool `url:"force,omitempty"`
}

func (bsrq GroupUpdateRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

// GroupUpdate updates existing Compute group within Basic Service and apply new settings to its computes as necessary
func (b BService) GroupUpdate(ctx context.Context, req GroupUpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupUpdate"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
