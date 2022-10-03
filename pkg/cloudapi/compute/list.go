package compute

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includedeleted,omitempty"`
	Page           uint64 `url:"page,omitempty"`
	Size           uint64 `url:"size,omitempty"`
}

func (c Compute) List(ctx context.Context, req ListRequest) (ComputeList, error) {

	url := "/cloudapi/compute/list"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
