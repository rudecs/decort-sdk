package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type TagAddRequest struct {
	ComputeId uint64 `url:"computeId"`
	Key       string `url:"key"`
	Value     string `url:"value"`
}

func (crq TagAddRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}

	if crq.Value == "" {
		return errors.New("validation-error: field Value can not be empty")
	}

	return nil
}

func (c Compute) TagAdd(ctx context.Context, req TagAddRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/tagAdd"
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
