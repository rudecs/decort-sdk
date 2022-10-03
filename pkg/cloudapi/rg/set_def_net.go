package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type SetDefNetRequest struct {
	RGID    uint64 `url:"rgId"`
	NetType string `url:"netType"`
	NetID   uint64 `url:"netId"`
	Reason  string `url:"reason,omitempty"`
}

func (rgrq SetDefNetRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	if !validators.StringInSlice(rgrq.NetType, []string{"PUBLIC", "PRIVATE"}) {
		return errors.New("field NetType can only be one of 'PUBLIC' or 'PRIVATE'")
	}

	return nil
}

func (r RG) SetDefNet(ctx context.Context, req SetDefNetRequest) (uint64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	url := "/cloudapi/rg/setDefNet"
	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(string(res), 10, 64)
}
