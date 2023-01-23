package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for updaate account
type UpdateRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Name of the account
	// Required: false
	Name string `url:"name,omitempty"`

	// Max size of memory in MB
	// Required: false
	MaxMemoryCapacity uint64 `url:"maxMemoryCapacity,omitempty"`

	// Max size of aggregated vdisks in GB
	// Required: false
	MaxVDiskCapacity uint64 `url:"maxVDiskCapacity,omitempty"`

	// Max number of CPU cores
	// Required: false
	MaxCPUCapacity uint64 `url:"maxCPUCapacity,omitempty"`

	// Max sent/received network transfer peering
	// Required: false
	MaxNetworkPeerTransfer uint64 `url:"maxNetworkPeerTransfer,omitempty"`

	// Max number of assigned public IPs
	// Required: false
	MaxNumPublicIP uint64 `url:"maxNumPublicIP,omitempty"`

	// If true send emails when a user is granted access to resources
	// Required: false
	SendAccessEmails bool `url:"sendAccessEmails,omitempty"`

	// Limit (positive) or disable (0) GPU resources
	// Required: false
	GPUUnits uint64 `url:"gpu_units,omitempty"`
}

func (arq UpdateRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

// Update updates an account name and resource types and limits
func (a Account) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/account/update"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
