package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request  struct for grant access to SEP
type AccessGrantRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Account ID to grant access to the specified SEP. If 0,
	// the SEP will be available for all accounts with no exceptions
	// Required: true
	AccountID uint64 `url:"account_id"`
}

func (srq AccessGrantRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}

	return nil
}

// AccessGrant grant access to SEP
func (s SEP) AccessGrant(ctx context.Context, req AccessGrantRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/accessGrant"

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
