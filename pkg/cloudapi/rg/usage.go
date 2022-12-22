package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get report of resource usage
type UsageRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (rgrq UsageRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// Usage gets report resource usage on the resource group
func (r RG) Usage(ctx context.Context, req UsageRequest) (*RecordResourceUsage, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/usage"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordResourceUsage{}
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
