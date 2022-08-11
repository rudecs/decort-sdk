package computeci

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includeDeleted,omitempty"`
	Page           uint64 `url:"page,omitempty"`
	Size           uint64 `url:"size,omitempty"`
}

func (c ComputeCI) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (ComputeCIList, error) {
	url := "/cloudapi/computeci/list"
	computeciListRaw, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	computeciList := ComputeCIList{}
	if err := json.Unmarshal(computeciListRaw, &computeciList); err != nil {
		return nil, err
	}

	return computeciList, nil
}
