package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type SnapshotRollbackRequest struct {
	DiskId    uint64 `url:"diskId"`
	Label     string `url:"label"`
	TimeStamp uint64 `url:"timestamp"`
}

func (drq SnapshotRollbackRequest) Validate() error {
	if drq.DiskId == 0 {
		return errors.New("validation-error: field DiskId can not be empty or equal to 0")
	}

	if drq.Label == "" && drq.TimeStamp == 0 {
		return errors.New("validation-error: field Label or field TimeStamp can not be empty")
	}

	return nil
}

func (d Disks) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/snapshotRollback"
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
