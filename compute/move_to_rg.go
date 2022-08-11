package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type MoveToRGRequest struct {
	ComputeId uint64 `url:"computeId"`
	RGID      uint64 `url:"rgId"`
	Name      string `url:"name,omitempty"`
	Autostart bool   `url:"autostart,omitempty"`
	ForceStop bool   `url:"forceStop,omitempty"`
}

func (crq MoveToRGRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) MoveToRG(ctx context.Context, req MoveToRGRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/moveToRg"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}
	return result, nil
}
