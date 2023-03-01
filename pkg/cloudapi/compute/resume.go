package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for resume compute
type ResumeRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq ResumeRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Resume resume Compute from paused state
func (c Compute) Resume(ctx context.Context, req ResumeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/resume"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
