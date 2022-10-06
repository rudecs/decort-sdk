package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type UsageRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq UsageRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) Usage(ctx context.Context, req UsageRequest) (*Reserved, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/usage"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	usage := Reserved{}

	err = json.Unmarshal(res, &usage)
	if err != nil {
		return nil, err
	}

	return &usage, nil
}
