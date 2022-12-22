package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove group compute
type GroupComputeRemoveRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId"`

	// ID of the Compute GROUP
	// Required: true
	CompGroupID uint64 `url:"compgroupId"`

	// ID of the Compute
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (bsrq GroupComputeRemoveRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}
	if bsrq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// GroupComputeRemove makes group compute remove of the Basic Service
func (b BService) GroupComputeRemove(ctx context.Context, req GroupComputeRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupComputeRemove"

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
