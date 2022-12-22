package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list port forward rules
type ListPFWRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`
}

func (rgrq ListPFWRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

// ListPFW gets list port forward rules for the specified resource group
func (r RG) ListPFW(ctx context.Context, req ListPFWRequest) (ListPFW, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/listPFW"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListPFW{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
