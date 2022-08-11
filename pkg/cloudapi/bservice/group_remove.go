package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GroupRemoveRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	CompGroupID uint64 `url:"compgroupId"`
}

func (bsrq GroupRemoveRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupRemove(ctx context.Context, req GroupRemoveRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupRemove"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err

	}

	return strconv.ParseBool(string(res))
}
