package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for attach network
type NetAttachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// Network type
	// 'EXTNET' for connect to external network directly
	// and 'VINS' for connect to ViNS
	// Required: true
	NetType string `url:"netType"`

	// Network ID for connect to
	// For EXTNET - external network ID
	// For VINS - VINS ID
	// Required: true
	NetID uint64 `url:"netId"`

	// Directly required IP address for new network interface
	// Required: true
	IPAddr string `url:"ipAddr,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq NetAttachRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.NetType == "" {
		return errors.New("validation-error: field NetType must be set")
	}
	validator := validators.StringInSlice(crq.NetType, []string{"EXTNET", "VINS"})
	if !validator {
		return errors.New("validation-error: field NetType can be only EXTNET or VINS")
	}
	if crq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// NetAttach attach network to compute and gets info about network
func (c Compute) NetAttach(ctx context.Context, req NetAttachRequest) (*RecordNetAttach, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/netAttach"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordNetAttach{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
