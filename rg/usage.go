package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type UsageRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq UsageRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) Usage(ctx context.Context, req UpdateRequest, options ...opts.DecortOpts) (*ResourceUsage, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/usage"
	usageRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	usage := &ResourceUsage{}
	if err := json.Unmarshal(usageRaw, usage); err != nil {
		return nil, err
	}

	return usage, nil
}
