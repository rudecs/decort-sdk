package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type GroupUpdateVINSRequest struct {
	ServiceID   uint64   `url:"serviceId"`
	CompGroupID uint64   `url:"compgroupId"`
	VINSes      []uint64 `url:"vinses,omitempty"`
}

func (bsrq GroupUpdateVINSRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupUpdateVINS(ctx context.Context, req GroupUpdateVINSRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupUpdateVins"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
