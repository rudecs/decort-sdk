package rg

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AccessRevokeRequest struct {
	RGID   uint64 `url:"rgId"`
	User   string `url:"user"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq AccessRevokeRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	if rgrq.User == "" {
		return errors.New("field User can not be empty")
	}

	return nil
}

func (r RG) AccessRevoke(ctx context.Context, req AccessRevokeRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/rg/accessRevoke"
	res, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
