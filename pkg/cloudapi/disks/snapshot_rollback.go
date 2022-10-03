package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type SnapshotRollbackRequest struct {
	DiskID    uint64 `url:"diskId"`
	Label     string `url:"label"`
	TimeStamp uint64 `url:"timestamp"`
}

func (drq SnapshotRollbackRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	if drq.Label == "" && drq.TimeStamp == 0 {
		return errors.New("validation-error: field Label or field TimeStamp can not be empty")
	}

	return nil
}

func (d Disks) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/snapshotRollback"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil

}
