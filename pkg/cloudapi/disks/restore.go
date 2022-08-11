package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type RestoreRequest struct {
	DiskId uint64 `url:"diskId"`
	Reason string `url:"reason"`
}

func (drq RestoreRequest) Validate() error {
	if drq.DiskId == 0 {
		return errors.New("validation-error: field DiskId can not be empty or equal to 0")
	}

	if drq.Reason == "" {
		return errors.New("validation-error: field Reason can not be empty")
	}

	return nil
}

func (d Disks) Restore(ctx context.Context, req RestoreRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/restore"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := d.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil

}
