package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create resource group
type CreateRequest struct {
	// Account, which will own this resource group
	// Required: true
	AccountID uint64 `url:"accountId"`

	// Grid ID
	// Required: true
	GID uint64 `url:"gid"`

	// Name of this resource group. Must be unique within the account
	// Required: true
	Name string `url:"name"`

	// Max size of memory in MB
	// Required: false
	MaxMemoryCapacity uint64 `url:"maxMemoryCapacity,omitempty"`

	// Max size of aggregated virtual disks in GB
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

	// Username - owner of this resource group.
	// Leave blank to set current user as owner
	// Required: false
	Owner string `url:"owner,omitempty"`

	// Type of the default network for this resource group.
	// virtual machines created in this resource group will be by default connected to this network.
	// Should be one of:
	//	- PRIVATE
	//	- PUBLIC
	//	- NONE
	// Required: false
	DefNet string `url:"def_net,omitempty"`

	// Private network IP CIDR if default network PRIVATE
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty"`

	// Text description of this resource group
	// Required: false
	Description string `url:"desc,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`

	// External network ID
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty"`

	// External IP address
	// Required: false
	ExtIP string `url:"extIp,omitempty"`

	// Register computes in registration system
	// Required: false
	RegisterComputes bool `url:"registerComputes,omitempty"`

	// List of strings with pools i.e.: ["sep1_poolName1", "sep2_poolName2"]
	// Required: false
	UniqPools []string `url:"unuqPools,omitempty"`

	// Resource types available to create in this account
	// Each element in a resource type slice should be one of:
	//	- compute
	//	- vins
	//	- k8s
	//	- openshift
	//	- lb
	//	- flipgroup
	// Required: false
	ResTypes []string `url:"resourceTypes,omitempty"`
}

func (rgrq CreateRequest) validate() error {
	if rgrq.AccountID == 0 {
		return errors.New("field AccountID can not be empty or equal to 0")
	}
	if rgrq.GID == 0 {
		return errors.New("field GID can not be empty or equal to 0")
	}
	if len(rgrq.Name) < 2 {
		return errors.New("field Name can not be shorter than two bytes")
	}
	if len(rgrq.ResTypes) > 0 {
		for _, value := range rgrq.ResTypes {
			validate := validators.StringInSlice(value, []string{"compute", "vins", "k8s", "openshift", "lb", "flipgroup"})
			if !validate {
				return errors.New("validation-error: Every resource type specified should be one of [compute, vins, k8s, openshift, lb, flipgroup]")
			}
		}
	}

	return nil
}

// Create creates resource group
func (r RG) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/rg/create"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
