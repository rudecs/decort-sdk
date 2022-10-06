package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	RGID        uint64 `url:"rgId"`
	Force       bool   `url:"force,omitempty"`
	Permanently bool   `url:"permanently,omitempty"`
	Reason      string `url:"reason,omitempty"`
}

func (rgrq DeleteRequest) Validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}

	return nil
}

func (r RG) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/delete"

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