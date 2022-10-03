package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ResizeRequest struct {
	DiskID uint64 `url:"diskId"`
	Size   uint64 `url:"size"`
}

func (drq ResizeRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}
	if drq.Size == 0 {
		return errors.New("validation-error: field Size must be set")
	}
	return nil
}

func (d Disks) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/resize"

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

func (d Disks) Resize2(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/resize2"

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
