package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RestoreRequest struct {
	DiskID uint64 `url:"diskId"`
	Reason string `url:"reason"`
}

func (drq RestoreRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

func (d Disks) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/restore"

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
