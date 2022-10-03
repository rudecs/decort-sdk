package computeci

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includeDeleted,omitempty"`
	Page           uint64 `url:"page,omitempty"`
	Size           uint64 `url:"size,omitempty"`
}

func (c ComputeCI) List(ctx context.Context, req ListRequest) (ComputeCIList, error) {
	url := "/cloudapi/computeci/list"
	computeciListRaw, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	computeciList := ComputeCIList{}
	if err := json.Unmarshal(computeciListRaw, &computeciList); err != nil {
		return nil, err
	}

	return computeciList, nil
}
