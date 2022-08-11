package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AffinityLabelSetRequest struct {
	ComputeId     uint64 `url:"computeId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityLabelSetRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

func (c Compute) AffinityLabelSet(ctx context.Context, req AffinityLabelSetRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/affinityLabelSet"
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
