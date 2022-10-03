package extnet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListComputesRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (erq ListComputesRequest) Validate() error {
	if erq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (e ExtNet) ListComputes(ctx context.Context, req ListComputesRequest) (ExtNetComputesList, error) {
	url := "/cloudapi/extnet/listComputes"

	extnetComputesListRaw, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	extnetComputesList := ExtNetComputesList{}
	err = json.Unmarshal(extnetComputesListRaw, &extnetComputesList)
	if err != nil {
		return nil, err
	}

	return extnetComputesList, nil

}
