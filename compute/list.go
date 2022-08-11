package compute

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includedeleted,omitempty"`
	Page           uint64 `url:"page,omitempty"`
	Size           uint64 `url:"size,omitempty"`
}

func (c Compute) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (ComputeList, error) {

	url := "/compute/list"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	computeList := ComputeList{}

	err = json.Unmarshal(res, &computeList)
	if err != nil {
		return nil, err
	}

	return computeList, nil
}
