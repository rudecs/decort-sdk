package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update account
type UpdateRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// Display name
	// Required: true
	Name string `url:"name" json:"name"`

	// Name of the account
	// Required: true
	Username string `url:"username" json:"username"`

	// Email
	// Required: false
	EmailAddress string `url:"emailaddress,omitempty" json:"emailaddress,omitempty"`

	// Max size of memory in MB
	// Required: false
	MaxMemoryCapacity uint64 `url:"maxMemoryCapacity,omitempty" json:"maxMemoryCapacity,omitempty"`

	// Max size of aggregated vdisks in GB
	// Required: false
	MaxVDiskCapacity uint64 `url:"maxVDiskCapacity,omitempty" json:"maxVDiskCapacity,omitempty"`

	// Max number of CPU cores
	// Required: false
	MaxCPUCapacity uint64 `url:"maxCPUCapacity,omitempty" json:"maxCPUCapacity,omitempty"`

	// Max sent/received network transfer peering
	// Required: false
	MaxNetworkPeerTransfer uint64 `url:"maxNetworkPeerTransfer,omitempty" json:"maxNetworkPeerTransfer,omitempty"`

	// Max number of assigned public IPs
	// Required: false
	MaxNumPublicIP uint64 `url:"maxNumPublicIP,omitempty" json:"maxNumPublicIP,omitempty"`

	// If true send emails when a user is granted access to resources
	// Required: false
	SendAccessEmails bool `url:"sendAccessEmails,omitempty" json:"sendAccessEmails,omitempty"`

	// Limit (positive) or disable (0) GPU resources
	// Required: false
	GPUUnits uint64 `url:"gpu_units,omitempty" json:"gpu_units,omitempty"`

	// List of strings with pools
	// i.e.: ["sep1_poolName1", "sep2_poolName2", etc]
	// Required: false
	UniqPools []string `url:"uniqPools,omitempty" json:"uniqPools,omitempty"`
}

func (arq UpdateRequest) validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// Update updates an account name and resource types and limits
func (a Account) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/update"

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
