package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RedeployRequest struct {
	ComputeID uint64 `url:"computeId"`
	ImageID   uint64 `url:"imageId,omitempty"`
	DiskSize  uint64 `url:"diskSize,omitempty"`
	DataDisks string `url:"dataDisks,omitempty"`
	AutoStart bool   `url:"autoStart,omitempty"`
	ForceStop bool   `url:"forceStop,omitempty"`
}

func (crq RedeployRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) Redeploy(ctx context.Context, req RedeployRequest) (bool, error) {
	err := req.Validate()
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
