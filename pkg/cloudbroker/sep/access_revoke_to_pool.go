package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for revoke access to pool SEP
type AccessRevokeToPoolRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`

	// Pool name
	// Required: true
	PoolName string `url:"pool_name" json:"pool_name"`

	// Account ID to grant access to the specified pool SEP
	// Required: false
	AccountID uint64 `url:"account_id,omitempty" json:"account_id,omitempty"`

	// Resource group ID to grant access to the specified pool SEP
	// Required: false
	RGID uint64 `url:"resgroup_id,omitempty" json:"resgroup_id,omitempty"`
}

func (srq AccessRevokeToPoolRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.PoolName == "" {
		return errors.New("validation-error: field PoolName must be set")
	}

	return nil
}

// AccessRevokeToPool revoke access to pool SEP
func (s SEP) AccessRevokeToPool(ctx context.Context, req AccessRevokeToPoolRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/accessRevokeToPool"

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
