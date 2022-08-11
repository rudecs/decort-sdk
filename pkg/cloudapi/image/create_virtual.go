package image

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateVirtualRequest struct {
	Name     string `url:"name"`
	TargetId uint64 `url:"targetId"`
}

func (irq CreateVirtualRequest) Validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if irq.TargetId == 0 {
		return errors.New("validation-error: field TargetId can not be empty or equal to 0")
	}

	return nil
}

func (i Image) CreateVirtual(ctx context.Context, req CreateVirtualRequest, options ...opts.DecortOpts) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/image/createVirtual"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := i.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
