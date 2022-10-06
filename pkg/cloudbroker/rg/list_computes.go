package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListComputesRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq ListComputesRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) ListComputes(ctx context.Context, req ListComputesRequest) (ListComputes, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/rg/listComputes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	listComputes := ListComputes{}

	err = json.Unmarshal(res, &listComputes)
	if err != nil {
		return nil, err
	}

	return listComputes, nil
}
