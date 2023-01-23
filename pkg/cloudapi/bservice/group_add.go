package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create new compute group within BasicService
type GroupAddRequest struct {
	// ID of the Basic Service to add a group to
	// Required: true
	ServiceID uint64 `url:"serviceId"`

	// Name of the Compute Group to add
	// Required: true
	Name string `url:"name"`

	// Computes number. Defines how many computes must be there in the group
	// Required: true
	Count uint64 `url:"count"`

	// Compute CPU number. All computes in the group have the same CPU count
	// Required: true
	CPU uint64 `url:"cpu"`

	// Compute RAM volume in MB. All computes in the group have the same RAM volume
	// Required: true
	RAM uint64 `url:"ram"`

	// Compute boot disk size in GB
	// Required: true
	Disk uint64 `url:"disk"`

	// OS image ID to create computes from
	// Required: true
	ImageID uint64 `url:"imageId"`

	// Compute driver
	// should be one of:
	//	- KVM_X86
	//	- KVM_PPC
	// Required: true
	Driver string `url:"driver"`

	// Storage endpoint provider ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty"`

	// Pool to use if sepId is set, can be also empty if needed to be chosen by system
	// Required: false
	SEPPool string `url:"sepPool,omitempty"`

	// Group role tag. Can be empty string, does not have to be unique
	// Required: false
	Role string `url:"role,omitempty"`

	// List of ViNSes to connect computes to
	// Required: false
	VINSes []uint64 `url:"vinses,omitempty"`

	// List of external networks to connect computes to
	// Required: false
	ExtNets []uint64 `url:"extnets,omitempty"`

	// Time of Compute Group readiness
	// Required: false
	TimeoutStart uint64 `url:"timeoutStart"`
}

func (bsrq GroupAddRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}
	if bsrq.Name == "" {
		return errors.New("field Name can not be empty")
	}
	if bsrq.Count == 0 {
		return errors.New("field Count can not be empty or equal to 0")
	}
	if bsrq.CPU == 0 {
		return errors.New("field CPU can not be empty or equal to 0")
	}
	if bsrq.RAM == 0 {
		return errors.New("field RAM can not be empty or equal to 0")
	}
	if bsrq.Disk == 0 {
		return errors.New("field Disk can not be empty or equal to 0")
	}
	if bsrq.ImageID == 0 {
		return errors.New("field ImageID can not be empty or equal to 0")
	}
	if bsrq.Driver == "" {
		return errors.New("field Driver can not be empty")
	}

	return nil
}

// GroupAdd creates new Compute Group within BasicService.
// Compute Group is NOT started automatically,
// so you need to explicitly start it
func (b BService) GroupAdd(ctx context.Context, req GroupAddRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/groupAdd"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
