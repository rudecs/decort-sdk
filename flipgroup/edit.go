package flipgroup

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type EditRequest struct {
	FlipGroupID uint64 `url:"flipgroupId"`
	Name        string `url:"name,omitempty"`
	Description string `url:"desc,omitempty"`
}

func (frq EditRequest) Validate() error {
	if frq.FlipGroupID == 0 {
		return errors.New("field FlipGroupID can not be empty or equal to 0")
	}

	return nil
}

func (f FlipGroup) Edit(ctx context.Context, req EditRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/edit"
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
