package rg

import (
	"context"
	"errors"
	"net/http"
)

type MassDisableRequest struct {
	RGIDs  []uint64 `url:"rgIds"`
	Reason string   `url:"reason,omitempty"`
}

func (rgrq MassDisableRequest) Validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

func (r RG) MassDisable(ctx context.Context, req MassDisableRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massDisable"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
