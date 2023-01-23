package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for unshare data disk
type UnshareRequest struct {
	// ID of the disk to unshare
	// Required: true
	DiskID uint64 `url:"diskId"`
}

func (drq UnshareRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

// Unshare unshares data disk
func (d Disks) Unshare(ctx context.Context, req UnshareRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/unshare"

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
