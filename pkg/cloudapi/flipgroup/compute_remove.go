package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for remove compute instance
type ComputeRemoveRequest struct {
	// ID of the Floating IP group to remove compute instance from
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId"`

	// ID of the compute instance to remove
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (frq ComputeRemoveRequest) validate() error {
	if frq.FLIPGroupID == 0 {
		return errors.New("field FLIPGroupID can not be empty or equal to 0")
	}
	if frq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// ComputeRemove remove compute instance from the Floating IP group
func (f FLIPGroup) ComputeRemove(ctx context.Context, req ComputeRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
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
