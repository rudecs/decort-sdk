package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type RenameRequest struct {
	DiskID uint64 `url:"diskId"`
	Name   string `url:"name"`
}

func (drq RenameRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	if drq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	return nil
}

func (d Disks) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/rename"

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
