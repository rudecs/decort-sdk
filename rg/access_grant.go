package rg

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type AccessGrantRequest struct {
	RGID   uint64 `url:"rgId"`
	User   string `url:"user"`
	Right  string `url:"right"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq AccessGrantRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	if rgrq.User == "" {
		return errors.New("field User can not be empty")
	}

	if !validators.StringInSlice(rgrq.Right, []string{"R", "RCX", "ARCXDU"}) {
		return errors.New("field Right can only be one of 'R', 'RCX' or 'ARCXDU'")
	}

	return nil
}

func (r RG) AccessGrant(ctx context.Context, req AccessGrantRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/rg/accessGrant"
	res, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
