package compute

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type PFWDelRequest struct {
	ComputeId       uint64 `url:"computeId"`
	PFWId           uint64 `url:"ruleId,omitempty"`
	PublicPortStart uint64 `url:"publicPortStart,omitempty"`
	PublicPortEnd   uint64 `url:"publicPortEnd,omitempty"`
	LocalBasePort   uint64 `url:"localBasePort,omitempty"`
	Proto           string `url:"proto,omitempty"`
}

func (crq PFWDelRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) PFWDel(ctx context.Context, req PFWDelRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/compute/pfwDel"
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
