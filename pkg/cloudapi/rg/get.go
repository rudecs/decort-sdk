package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get detailed information about resource group
type GetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq GetRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// Get gets current configuration of the resource group
func (r RG) Get(ctx context.Context, req GetRequest) (*RecordResourceGroup, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/get"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordResourceGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
