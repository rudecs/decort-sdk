package compute

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type PFWListRequest struct {
	ComputeId uint64 `url:"computeId"`
}

func (crq PFWListRequest) Validate() error {
	if crq.ComputeId == 0 {
		return errors.New("validation-error: field ComputeId can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) PFWList(ctx context.Context, req PFWListRequest, options ...opts.DecortOpts) (PFWList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/compute/pfwList"
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
		return nil, err
	}

	pfwList := PFWList{}

	err = json.Unmarshal(res, &pfwList)
	if err != nil {
		return nil, err
	}

	return pfwList, nil
}
