package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	DiskID      uint64 `url:"diskId"`
	Detach      bool   `url:"detach,omitempty"`
	Permanently bool   `url:"permanently,omitempty"`
	Reason      string `url:"reason,omitempty"`
}

func (drq DeleteRequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

func (d Disks) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/delete"

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