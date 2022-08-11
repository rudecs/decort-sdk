package compute

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateTemplateRequest struct {
	ComputeId uint64 `url:"computeId"`
	Name      string `url:"name"`
	Async     bool   `url:"async"`
}

func (crq CreateTemplateRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	if crq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	return nil
}

func (c Compute) CreateTemplate(ctx context.Context, req CreateTemplateRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	req.Async = false

	url := "/compute/createTemplate"
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
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, nil
	}

	return result, nil
}

func (c Compute) CreateTemplateAsync(ctx context.Context, req CreateTemplateRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	req.Async = true

	url := "/compute/createTemplate"
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

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
