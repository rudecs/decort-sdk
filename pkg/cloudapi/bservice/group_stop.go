package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for stop the specified Compute Group
type GroupStopRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// ID of the Compute Group to stop
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId"`

	// Force stop Compute Group
	// Required: true
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

func (bsrq GroupStopRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

// GroupStop stops the specified Compute Group within BasicService
func (b BService) GroupStop(ctx context.Context, req GroupStopRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupStop"

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
