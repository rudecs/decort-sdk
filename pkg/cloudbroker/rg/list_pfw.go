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
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) ListPFW(ctx context.Context, req ListPFWRequest) (ListPFW, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/listPFW"

	pfwListRaw, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	pfwList := ListPFW{}

	err = json.Unmarshal(pfwListRaw, &pfwList)
	if err != nil {
		return nil, err
	}

	return pfwList, nil
}
