package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GroupUpdateRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	CompGroupID uint64 `url:"compgroupId"`
	Name        string `url:"name,omitempty"`
	Role        string `url:"role,omitempty"`
	CPU         uint64 `url:"cpu,omitempty"`
	RAM         uint64 `url:"ram,omitempty"`
	Disk        uint64 `url:"disk,omitempty"`
	Force       bool   `url:"force,omitempty"`
}

func (bsrq GroupUpdateRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupUpdate(ctx context.Context, req GroupUpdateRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/groupUpdate"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
