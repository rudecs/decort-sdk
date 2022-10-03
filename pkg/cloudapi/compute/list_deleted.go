package compute

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (c Compute) ListDeleted(ctx context.Context, req ListDeletedRequest) (ComputeList, error) {

	url := "/cloudapi/compute/listDeleted"

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
