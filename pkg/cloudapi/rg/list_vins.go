package rg

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListVINSRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq ListVINSRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) ListVINS(ctx context.Context, req ListVINSRequest, options ...opts.DecortOpts) (VINSList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listVins"
	vinsListRaw, err := r.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	vinsList := VINSList{}
	if err := json.Unmarshal(vinsListRaw, &vinsList); err != nil {
		return nil, err
	}

	return vinsList, nil
}
