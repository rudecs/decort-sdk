package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for enable resource group
type EnableRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq EnableRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// Enable enables resource group
func (r RG) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/rg/enable"

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
