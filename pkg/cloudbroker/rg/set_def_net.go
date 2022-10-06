package rg

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type SetDefNetRequest struct {
	RGID    uint64 `url:"rgId"`
	NetType string `url:"netType"`
	NetID   uint64 `url:"netId,omitempty"`
	Reason  string `url:"reason,omitempty"`
}

func (rgrq SetDefNetRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	validate := validators.StringInSlice(rgrq.NetType, []string{"PUBLIC", "PRIVATE"})
	if !validate {
		return errors.New("validation-error: field NetType must be one of PRIVATE or PUBLIC")
	}

	return nil
}

func (r RG) SetDefNet(ctx context.Context, req SetDefNetRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/rg/setDefNet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
