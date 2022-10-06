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
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) ListVINS(ctx context.Context, req ListVINSRequest) (ListVINS, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/listVins"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	listVINS := ListVINS{}

	err = json.Unmarshal(res, &listVINS)
	if err != nil {
		return nil, err
	}

	return listVINS, nil
}
