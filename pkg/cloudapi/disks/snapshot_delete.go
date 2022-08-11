package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotDeleteRequest struct {
	DiskId uint64 `url:"diskId"`
	Label  string `url:"label"`
}

func (drq SnapshotDeleteRequest) Validate() error {
	if drq.DiskId == 0 {
		return errors.New("validation-error: field DiskId can not be empty or equal to 0")
	}

	if drq.Label == "" {
		return errors.New("validation-error: field Label can not be empty")
	}

	return nil
}

func (d Disks) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/snapshotDelete"
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
