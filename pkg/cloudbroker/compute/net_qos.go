package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for update QOS
type NetQOSRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Network ID
	// Required: true
	NetID uint64 `url:"netId" json:"netId"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	// Required: true
	NetType string `url:"netType" json:"netType"`

	// Internal traffic, kbit
	// Required: false
	IngressRate uint64 `url:"ingress_rate,omitempty" json:"ingress_rate,omitempty"`

	// Internal traffic burst, kbit
	// Required: false
	IngressBurst uint64 `url:"ingress_burst,omitempty" json:"ingress_burst,omitempty"`

	// External traffic rate, kbit
	// Required: false
	EgressRate uint64 `url:"egress_rate,omitempty" json:"egress_rate,omitempty"`
}

func (crq NetQOSRequest) validate() error {
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

// NetQOS update compute interfaces QOS
func (c Compute) NetQOS(ctx context.Context, req NetQOSRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/netQos"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
