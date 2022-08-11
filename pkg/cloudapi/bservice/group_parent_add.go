package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GroupParentAddRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	CompGroupID uint64 `url:"compgroupId"`
	ParentID    uint64 `url:"parentId"`
}

func (bsrq GroupParentAddRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	if bsrq.ParentID == 0 {
		return errors.New("field ParentID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupParentAdd(ctx context.Context, req GroupParentAddRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupParentAdd"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
