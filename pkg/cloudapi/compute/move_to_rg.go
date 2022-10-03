package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type MoveToRGRequest struct {
	ComputeID uint64 `url:"computeId"`
	RGID      uint64 `url:"rgId"`
	Name      string `url:"name,omitempty"`
	Autostart bool   `url:"autostart,omitempty"`
	ForceStop bool   `url:"forceStop,omitempty"`
}

func (crq MoveToRGRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) MoveToRG(ctx context.Context, req MoveToRGRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/moveToRg"

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
