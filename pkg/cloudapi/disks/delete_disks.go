package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DisksDeleteRequest struct {
	DisksIDs    []uint64 `url:"diskIds"`
	Reason      string   `url:"reason"`
	Permanently bool     `url:"permanently"`
}

func (drq DisksDeleteRequest) Validate() error {
	if len(drq.DisksIDs) == 0 {
		return errors.New("validation-error: field DisksIDs must include one or more disks ids")
	}

	return nil
}

func (d Disks) DeleteDisks(ctx context.Context, req DisksDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/deleteDisks"

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
