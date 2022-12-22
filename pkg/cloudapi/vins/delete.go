package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete VINS
type DeleteRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Set to True if you want force delete non-empty VINS.
	// Primarily, VINS is considered non-empty if it has virtual machines connected to it,
	// and force flag will detach them from the VINS being deleted.
	// Otherwise method will return an error
	// Required: false
	Force bool `url:"force,omitempty"`

	// Set to True if you want to destroy VINS and all linked resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be restored later within the recycle bin's purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty"`
}

func (vrq DeleteRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// Delete deletes VINS
func (v VINS) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/delete"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
