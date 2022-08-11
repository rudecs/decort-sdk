package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotCreateRequest struct {
	ServiceID uint64 `url:"serviceId"`
	Label     string `url:"label"`
}

func (bsrq SnapshotCreateRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.Label == "" {
		return errors.New("field Label can not be empty")
	}

	return nil
}

func (b BService) SnapshotCreate(ctx context.Context, req SnapshotCreateRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/bservice/snapshotCreate"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
