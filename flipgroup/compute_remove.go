package flipgroup

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ComputeRemoveRequest struct {
	FlipGroupID uint64 `url:"flipgroupId"`
	ComputeID   uint64 `url:"computeId"`
}

func (frq ComputeRemoveRequest) Validate() error {
	if frq.FlipGroupID == 0 {
		return errors.New("field FlipGroupID can not be empty or equal to 0")
	}
	if frq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (f FlipGroup) ComputeRemove(ctx context.Context, req ComputeRemoveRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/computeRemove"
	res, err := f.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, nil
	}

	return result, nil
}
