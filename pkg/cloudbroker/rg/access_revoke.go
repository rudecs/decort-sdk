package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type AccessRevokeRequest struct {
	RGID   uint64 `url:"rgId"`
	User   string `url:"user"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq AccessRevokeRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	if rgrq.User == "" {
		return errors.New("validation-error: field User must be set")
	}

	return nil
}

func (r RG) AccessRevoke(ctx context.Context, req AccessRevokeRequest) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/accessRevoke"

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
