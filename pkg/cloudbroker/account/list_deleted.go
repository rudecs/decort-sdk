package account

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListDeletedRequest struct {
	Page uint64 `url:"page,omitempty"`
	Size uint64 `url:"size,omitempty"`
}

func (a Account) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListInfoResponse, error) {
	url := "/cloudbroker/account/listDeleted"

	result := ListInfoResponse{}
	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return ListInfoResponse{}, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return ListInfoResponse{}, err
	}

	return result, err
}
