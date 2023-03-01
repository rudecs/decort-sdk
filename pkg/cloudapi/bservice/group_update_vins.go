package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update VINS settings
type GroupUpdateVINSRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId"`

	// List of ViNSes to connect computes
	// Required: false
	VINSes []uint64 `url:"vinses,omitempty" json:"vinses,omitempty"`
}

func (bsrq GroupUpdateVINSRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

// GroupUpdateVINS update ViNS settings for the group according to the new list
func (b BService) GroupUpdateVINS(ctx context.Context, req GroupUpdateVINSRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupUpdateVins"

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
