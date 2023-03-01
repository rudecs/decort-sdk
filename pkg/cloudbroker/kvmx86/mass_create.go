package kvmx86

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for mass create KVM x86
type MassCreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name" json:"name"`

	// Number of VMs
	// Required: true
	Count uint64 `url:"count" json:"count"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu" json:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram" json:"ram"`

	// Image ID
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Size of the boot disk in GB
	// Required: true
	BootDisk uint64 `url:"bootDisk" json:"bootDisk"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: true
	SEPID uint64 `url:"sepId" json:"sepId"`

	// Pool to use if SEP ID is set, can be also empty if needed to be chosen by system
	// Required: true
	Pool string `url:"pool" json:"pool"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty" json:"netType,omitempty"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty" json:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// Input data for cloud-init facility
	// Required: false
	UserData string `url:"userdata,omitempty" json:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start after create of not
	// Required: false
	Start bool `url:"start,omitempty" json:"start,omitempty"`

	// Reason to action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (krq MassCreateRequest) validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.Count == 0 {
		return errors.New("validation-error: field Count must be set")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU must be set")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM must be set")
	}
	if krq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// MassCreate creates KVM x86 computes based on specified OS image
func (k KVMX86) MassCreate(ctx context.Context, req MassCreateRequest) ([]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/kvmx86/massCreate"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	computes := make([]uint64, 0)

	err = json.Unmarshal(res, &computes)
	if err != nil {
		return nil, err
	}

	return computes, nil
}
