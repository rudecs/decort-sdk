package disks

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type GetRequest struct {
	DiskId uint64 `url:"diskId"`
}

func (drq GetRequest) Validate() error {
	if drq.DiskId == 0 {
		return errors.New("validation-error: field DiskId can not be empty or equal to 0")
	}

	return nil
}

func (d Disks) Get(ctx context.Context, req GetRequest, options ...opts.DecortOpts) (*DiskRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/disks/get"
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
		return nil, err
	}

	disk := &DiskRecord{}

	err = json.Unmarshal(res, disk)
	if err != nil {
		return nil, err
	}

	return disk, nil

}
