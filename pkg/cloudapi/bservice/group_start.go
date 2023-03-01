package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for start the specified Compute Group
type GroupStartRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// ID of the Compute Group to start
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId"`
}

func (bsrq GroupStartRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

// GroupStart starts the specified Compute Group within BasicService
func (b BService) GroupStart(ctx context.Context, req GroupStartRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupStart"

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
