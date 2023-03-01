package kvmppc

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create KVM PowerPC VM from scratch
type CreateBlankRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name" json:"name"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu" json:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram" json:"ram"`

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

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

func (krq CreateBlankRequest) validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU can not be empty or equal to 0")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM can not be empty or equal to 0")
	}
	if krq.BootDisk == 0 {
		return errors.New("validation-error: field BootDisk can not be empty or equal to 0")
	}
	if krq.SEPID == 0 {
		return errors.New("validation-error: field SepID can not be empty or equal to 0")
	}
	if krq.Pool == "" {
		return errors.New("validation-error: field Pool can not be empty")
	}

	return nil
}

// CreateBlank creates KVM PowerPC VM from scratch
func (k KVMPPC) CreateBlank(ctx context.Context, req CreateBlankRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/kvmppc/createBlank"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
