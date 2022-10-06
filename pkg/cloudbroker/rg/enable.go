package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type EnableRequest struct {
	RGID   uint64 `url:"rgId"`
	Reason string `url:"reason"`
}

func (rgrq EnableRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/enable"

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