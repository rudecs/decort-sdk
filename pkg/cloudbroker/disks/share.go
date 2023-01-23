package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for share data disk
type ShareRequest struct {
	// ID of the disk to share
	// Required: true
	DiskID uint64 `url:"diskId"`
}

func (drq ShareRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

// Share shares data disk
func (d Disks) Share(ctx context.Context, req ShareRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/share"

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
