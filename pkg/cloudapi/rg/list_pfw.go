package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListPFWRequest struct {
	RGID uint64 `url:"rgId"`
}

func (rgrq ListPFWRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (r RG) ListPFW(ctx context.Context, req ListPFWRequest) (PortForwardList, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/listPFW"
	pfwListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	pfwList := PortForwardList{}
	if err := json.Unmarshal(pfwListRaw, &pfwList); err != nil {
		return nil, err
	}

	return pfwList, nil
}
