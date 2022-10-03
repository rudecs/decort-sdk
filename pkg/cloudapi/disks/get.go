package disks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	DiskID uint64 `url:"diskId"`
}

func (drq GetRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

func (d Disks) Get(ctx context.Context, req GetRequest) (*DiskRecord, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/disks/get"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
