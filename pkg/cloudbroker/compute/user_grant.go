package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for grant access to compute
type UserGrantRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Name of the user to add
	// Required: true
	Username string `url:"userName" json:"userName"`

	// Access type
	// Should be one of:
	//	- 'R' for Read only
	//	- 'RCX' for Write
	//	- 'ARCXDU' for Admin
	// Required: true
	AccessType string `url:"accesstype" json:"accesstype"`
}

func (crq UserGrantRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.Username == "" {
		return errors.New("validation-error: field Username must be set")
	}
	if crq.AccessType == "" {
		return errors.New("validation-error: field AccessType must be set")
	}
	validator := validators.StringInSlice(crq.AccessType, []string{"R", "RCX", "ARCXDU"})
	if !validator {
		return errors.New("validation-error: field AccessType can be only R, RCX or ARCXDU")
	}

	return nil
}

// UserGrant grant user access to the compute
func (c Compute) UserGrant(ctx context.Context, req UserGrantRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/userGrant"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
