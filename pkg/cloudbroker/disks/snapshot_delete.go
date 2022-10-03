package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type SnapshotDeleteRequest struct {
	DiskID uint64 `url:"diskId"`
	Label  string `url:"label"`
}

func (drq SnapshotDeleteRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Label == "" {
		return errors.New("validation-error: field Label must be set")
	}

	return nil
}

func (d Disks) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/snapshotDelete"

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
