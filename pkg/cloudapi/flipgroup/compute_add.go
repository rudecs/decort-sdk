package flipgroup

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add compute instance
type ComputeAddRequest struct {
	// ID of the Floating IP group to add compute instance to
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId"`

	// ID of the compute instance to add to this group
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (frq ComputeAddRequest) validate() error {
	if frq.FLIPGroupID == 0 {
		return errors.New("field FLIPGroupID can not be empty or equal to 0")
	}
	if frq.ComputeID == 0 {
		return errors.New("field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// ComputeAdd add compute instance to the Floating IP group
func (f FLIPGroup) ComputeAdd(ctx context.Context, req ComputeAddRequest) (bool, error) {
	err := req.validate()
	if err != nil {
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
