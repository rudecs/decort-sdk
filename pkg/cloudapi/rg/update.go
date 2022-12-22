package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for update resource group
type UpdateRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`

	// New name
	// Required: false
	Name string `url:"name,omitempty"`

	// New description
	// Required: false
	Description string `url:"desc,omitempty"`

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

	// Register computes in registration system
	// Required: false
	RegisterComputes bool `url:"registerComputes,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`

	// Resource types available to create in this account
	// Each element in a resource type slice must be one of:
	//	- compute
	//	- vins
	//	- k8s
	//	- openshift
	//	- lb
	//	- flipgroup
	// Required: false
	ResTypes []string `url:"resourceTypes,omitempty"`
}

func (rgrq UpdateRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
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

// Update updates resource group
func (r RG) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/rg/update"

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
