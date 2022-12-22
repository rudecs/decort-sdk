package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete resource group
type DeleteRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Set to True if you want force delete non-empty resource group
	// Required: false
	Force bool `url:"force,omitempty"`

	// Set to True if you want to destroy resource group and all linked resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be restored later within recycle bin's purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq DeleteRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

// Delete deletes resource group
func (r RG) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/delete"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
