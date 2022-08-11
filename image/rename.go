package image

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type RenameRequest struct {
	ImageId uint64 `url:"imageId"`
	Name    string `url:"name"`
}

func (irq RenameRequest) Validate() error {
	if irq.ImageId == 0 {
		return errors.New("validation-error: field ImageId can not be empty or equal to 0")
	}

	if irq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	return nil
}

func (i Image) Rename(ctx context.Context, req RenameRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/image/rename"
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
