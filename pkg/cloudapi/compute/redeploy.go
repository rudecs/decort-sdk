package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for redeploy
type RedeployRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// ID of the new OS image, if image change is required
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`

	// new size for the boot disk in GB, if boot disk size change is required
	// Required: false
	DiskSize uint64 `url:"diskSize,omitempty" json:"diskSize,omitempty"`

	// How to handle data disks connected to this compute instance,
	// KEEP, DETACH, DESTROY
	// Required: false
	DataDisks string `url:"dataDisks,omitempty" json:"dataDisks,omitempty"`

	// Should the compute be restarted upon successful redeploy
	// Required: false
	AutoStart bool `url:"autoStart,omitempty" json:"autoStart,omitempty"`

	// Set this flag to True to force stop running compute instance and redeploy next
	// Required: false
	ForceStop bool `url:"forceStop,omitempty" json:"forceStop,omitempty"`
}

func (crq RedeployRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// Redeploy redeploy compute
func (c Compute) Redeploy(ctx context.Context, req RedeployRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/redeploy"

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
