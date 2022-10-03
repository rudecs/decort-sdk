package account

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (a Account) List(ctx context.Context, req ListRequest) (ListInfoResponse, error) {
	url := "/cloudbroker/account/list"

	result := ListInfoResponse{}
	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListInfoResponse{}, err
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListInfoResponse{}, err
	}

	return result, nil
}
