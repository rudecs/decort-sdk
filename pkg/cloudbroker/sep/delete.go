package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete SEP
type DeleteRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`
}

func (srq DeleteRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// Delete deletes SEP by ID
func (s SEP) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/delete"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
