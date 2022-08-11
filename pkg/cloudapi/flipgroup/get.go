package flipgroup

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	FlipGroupID uint64 `url:"flipgroupId"`
}

func (frq GetRequest) Validate() error {
	if frq.FlipGroupID == 0 {
		return errors.New("field FlipGroupID can not be empty or equal to 0")
	}

	return nil
}

func (f FlipGroup) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*FlipGroupItem, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/flipgroup/get"
	res, err := f.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	fg := &FlipGroupItem{}
	err = json.Unmarshal(res, fg)
	if err != nil {
		return nil, err
	}

	return fg, nil
}
