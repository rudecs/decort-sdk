package rg

import (
	"context"
	"errors"
	"net/http"
)

type MassDeleteRequest struct {
	RGIDs       []uint64 `url:"rgIds"`
	Force       bool     `url:"force,omitempty"`
	Permanently bool     `url:"permanently,omitempty"`
	Reason      string   `url:"reason,omitempty"`
}

func (rgrq MassDeleteRequest) Validate() error {
	if len(rgrq.RGIDs) == 0 {
		return errors.New("validation-error: field RGIDs must be set")
	}

	return nil
}

func (r RG) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/massDelete"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
