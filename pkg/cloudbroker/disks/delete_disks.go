package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteDisksRequest struct {
	DiskIDS     []uint64 `url:"diskIds"`
	Reason      string   `url:"reason"`
	Permanently bool     `url:"permanently,omitempty"`
}

func (drq DeleteDisksRequest) Validate() error {
	if len(drq.DiskIDS) == 0 {
		return errors.New("validation-error: field DiskIDs must be set")
	}
	if drq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

func (d Disks) DeleteDisks(ctx context.Context, req DeleteDisksRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/deleteDisks"

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
