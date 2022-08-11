package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteRequest struct {
	DiskId      uint64 `url:"diskId"`
	Detach      bool   `url:"detach,omitempty"`
	Permanently bool   `url:"permanently,omitempty"`
	Reason      string `url:"reason,omitempty"`
}

func (d DeleteRequest) Validate() error {

	if d.DiskId == 0 {
		return errors.New("validation-error: field DiskId must be set")
	}
	return nil
}

func (d Disks) Delete(ctx context.Context, req DeleteRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/delete"
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
