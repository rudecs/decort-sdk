package rg

import (
	"context"
	"errors"
	"github.com/rudecs/decort-sdk/internal/validators"
	"net/http"
	"strconv"
)

type AccessGrantRequest struct {
	RGID   uint64 `url:"rgId"`
	User   string `url:"user"`
	Right  string `url:"right"`
	Reason string `url:"reason,omitempty"`
}

func (rgrq AccessGrantRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	if rgrq.User == "" {
		return errors.New("validation-error: field User must be set")
	}

	validate := validators.StringInSlice(rgrq.Right, []string{"R", "RCX", "ARCXDU"})
	if !validate {
		return errors.New("field Right can only be one of 'R', 'RCX' or 'ARCXDU'")
	}

	return nil
}

func (r RG) AccessGrant(ctx context.Context, req AccessGrantRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/accessGrant"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
