package compute

import (
	"context"
	"errors"
	"strings"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetConsoleURLRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq GetConsoleURLRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) GetConsoleURL(ctx context.Context, req GetConsoleURLRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/compute/getConsoleUrl"
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
