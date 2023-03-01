package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete compute
type DeleteRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Delete permanently
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Set True if you want to detach data disks (if any) from the compute before its deletion
	// Required: false
	DetachDisks bool `url:"detachDisks,omitempty" json:"detachDisks,omitempty"`
}

func (crq DeleteRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Delete deletes compute
func (c Compute) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/delete"

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
