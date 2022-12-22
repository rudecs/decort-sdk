package kvmx86

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create KVM x86 VM from scratch
type CreateBlankRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram"`

	// Size of the boot disk in GB
	// Required: true
	BootDisk uint64 `url:"bootDisk"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: true
	SEPID uint64 `url:"sepId"`

	// Pool to use if sepId is set, can be also empty if needed to be chosen by system
	// Required: true
	Pool string `url:"pool"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty"`
}

func (krq CreateBlankRequest) validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU must be set")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM must be set")
	}
	if krq.BootDisk == 0 {
		return errors.New("validation-error: field BootDisk must be set")
	}
	if krq.SEPID == 0 {
		return errors.New("validation-error: field SepID must be set")
	}
	if krq.Pool == "" {
		return errors.New("validation-error: field Pool must be set")
	}

	return nil
}

// CreateBlank creates KVM x86 VM from scratch
func (k KVMX86) CreateBlank(ctx context.Context, req CreateBlankRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/kvmx86/createBlank"

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
