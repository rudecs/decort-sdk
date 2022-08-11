package rg

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type DeleteRequest struct {
	RGID        uint64 `url:"rgId"`
	Force       bool   `url:"force,omitempty"`
	Permanently bool   `url:"permanently,omitempty"`
	Reason      string `url:"reason,omitempty"`
}

func (rgrq DeleteRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) Delete(ctx context.Context, req DeleteRequest, options ...opts.DecortOpts) (bool, error) {
	if err := req.Validate(); err != nil {
		return false, err
	}

	url := "/cloudapi/rg/delete"
	res, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(res))
}
