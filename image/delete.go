package image

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteRequest struct {
	ImageId     uint64 `url:"imageId"`
	Permanently bool   `url:"permanently"`
}

func (irq DeleteRequest) Validate() error {
	if irq.ImageId == 0 {
		return errors.New("validation-error: field ImageId can not be empty or equal to 0")
	}

	return nil
}

func (i Image) Delete(ctx context.Context, req DeleteRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/image/delete"
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
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil

}
