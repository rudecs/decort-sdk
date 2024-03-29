package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for grant access to resource group
type AccessGrantRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// User or group name to grant access
	// Required: true
	User string `url:"user" json:"user"`

	// Access rights to set,
	// Should be one of:
	//	- "R"
	//	- "RCX"
	//	- "ARCXDU"
	// Required: true
	Right string `url:"right" json:"right"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (rgrq AccessGrantRequest) validate() error {
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

// AccessGrant grants user or group access to the resource group as  specified
func (r RG) AccessGrant(ctx context.Context, req AccessGrantRequest) (bool, error) {
	err := req.validate()
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
