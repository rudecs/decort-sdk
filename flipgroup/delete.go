package flipgroup

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteRequest struct {
	FlipGroupID uint64 `url:"flipgroupId"`
}

func (frq DeleteRequest) Validate() error {
	if frq.FlipGroupID == 0 {
		return errors.New("field FlipGroupID can not be empty or equal to 0")
	}

	return nil
}

func (f FlipGroup) Delete(ctx context.Context, req DeleteRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/delete"
	res, err := f.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
