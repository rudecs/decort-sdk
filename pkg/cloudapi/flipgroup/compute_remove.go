package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
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

func (f FlipGroup) ComputeRemove(ctx context.Context, req ComputeRemoveRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/flipgroup/computeRemove"
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
