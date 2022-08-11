package compute

import (
	"context"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetLogRequest struct {
	ComputeId uint64 `url:"computeId"`
	Path      string `url:"path"`
}

func (crq GetLogRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	if crq.Path == "" {
		return errors.New("validation-error: field Path can not be empty")
	}

	return nil
}

func (c Compute) GetLog(ctx context.Context, req GetLogRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/compute/getLog"
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
		return "", err
	}

	return string(res), nil
}

func (c Compute) GetLogGet(ctx context.Context, req GetLogRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/compute/getLog"
	prefix := "/cloudapi"

	option := opts.New(options)
	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := c.client.DecortApiCall(ctx, typed.GET, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
