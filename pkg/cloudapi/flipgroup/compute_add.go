package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ComputeAddRequest struct {
	FlipGroupID uint64 `url:"flipgroupId"`
	ComputeID   uint64 `url:"computeId"`
}

func (frq ComputeAddRequest) Validate() error {
	if frq.FlipGroupID == 0 {
		return errors.New("field FlipGroupID can not be empty or equal to 0")
	}
	if frq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (f FlipGroup) ComputeAdd(ctx context.Context, req ComputeAddRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/computeAdd"
	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
