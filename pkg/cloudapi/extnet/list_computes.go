package extnet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list computes
type ListComputesRequest struct {
	// Filter by account ID
	// Required: true
	AccountID uint64 `url:"accountId"`
}

func (erq ListComputesRequest) validate() error {
	if erq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// ListComputes gets computes from account with extnets
func (e ExtNet) ListComputes(ctx context.Context, req ListComputesRequest) (ListExtNetComputes, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/extnet/listComputes"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListExtNetComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
