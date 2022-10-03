package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type GroupUpdateExtNetRequest struct {
	ServiceID   uint64   `url:"serviceId"`
	CompGroupID uint64   `url:"compgroupId"`
	ExtNets     []uint64 `url:"extnets,omitempty"`
}

func (bsrq GroupUpdateExtNetRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupUpdateExtNet(ctx context.Context, req GroupUpdateExtNetRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupUpdateExtnet"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
