package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for revoke access to SEP
type AccessRevokeRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Account ID to revoke access to the specified SEP
	// Required: true
	AccountID uint64 `url:"account_id"`
}

func (srq AccessRevokeRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// AccessRevoke revoke access to SEP
func (s SEP) AccessRevoke(ctx context.Context, req AccessRevokeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/accessRevoke"

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
