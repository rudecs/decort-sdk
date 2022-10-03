package extnet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetRequest struct {
	NetID uint64 `url:"net_id"`
}

func (erq GetRequest) Validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID can not be empty or equal to 0")
	}
	return nil
}

func (e ExtNet) Get(ctx context.Context, req GetRequest) (*ExtNetDetailed, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/extnet/get"

	extnetRaw, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	extnet := &ExtNetDetailed{}
	err = json.Unmarshal(extnetRaw, &extnet)
	if err != nil {
		return nil, err
	}

	return extnet, nil

}
