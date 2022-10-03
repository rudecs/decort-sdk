package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

func (r RG) ListVINS(ctx context.Context, req ListVINSRequest) (VINSList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listVins"
	VINSListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	VINSList := VINSList{}
	if err := json.Unmarshal(VINSListRaw, &VINSList); err != nil {
		return nil, err
	}

	return VINSList, nil
}
