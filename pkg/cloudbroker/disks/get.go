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
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

func (d Disks) Get(ctx context.Context, req GetRequest) (*Disk, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/disks/get"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	disk := Disk{}

	err = json.Unmarshal(res, &disk)
	if err != nil {
		return nil, err
	}

	return &disk, nil
}
