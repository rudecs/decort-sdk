package rg

import (
	"context"
	"errors"
	"net/http"
)

type MassEnableRequest struct {
	RGIDs  []uint64 `url:"rgIds"`
	Reason string   `url:"reason,omitempty"`
}

func (rgrq MassEnableRequest) Validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

func (r RG) MassEnable(ctx context.Context, req MassEnableRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massEnable"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
