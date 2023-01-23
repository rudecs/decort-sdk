package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for creating account
type CreateRequest struct {
	// Display name
	// Required: true
	Name string `url:"name"`

	// Name of the account
	// Required: true
	Username string `url:"username"`

	// Email
	// Required: false
	EmailAddress string `url:"emailaddress,omitempty"`

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

	// List of strings with pools
	// i.e.: ["sep1_poolName1", "sep2_poolName2", etc]
	// Required: false
	UniqPools []string `url:"uniqPools,omitempty"`
}

func (arq CreateRequest) validate() error {
	if arq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if arq.Username == "" {
		return errors.New("validation-error: field Username can not be empty")
	}

	return nil
}

// Create creates account
// Setting a cloud unit maximum to -1 or empty will not put any restrictions on the resource
func (a Account) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/account/create"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
