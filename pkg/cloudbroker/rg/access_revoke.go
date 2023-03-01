package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for revoke access
type AccessRevokeRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// User or group name to revoke access
	// Required: true
	User string `url:"user" json:"user"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (rgrq AccessRevokeRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if rgrq.User == "" {
		return errors.New("validation-error: field User must be set")
	}

	return nil
}

// AccessRevoke revokes specified user or group access from the resource group
func (r RG) AccessRevoke(ctx context.Context, req AccessRevokeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/rg/accessRevoke"

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
