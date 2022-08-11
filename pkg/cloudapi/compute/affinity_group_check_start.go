package compute

import (
	"context"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AffinityGroupCheckStartRequest struct {
	RGID          uint64 `url:"rgId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityGroupCheckStartRequest) Validate() error {
	if crq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if crq.AffinityLabel == "" {
		return errors.New("validation-error: field AffinityLabel can not be empty")
	}

	return nil
}

func (c Compute) AffinityGroupCheckStart(ctx context.Context, req AffinityGroupCheckStartRequest, options ...opts.DecortOpts) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/compute/affinityGroupCheckStart"
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
