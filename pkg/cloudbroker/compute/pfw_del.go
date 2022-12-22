package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete port forward rule
type PFWDelRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// ID of the rule to delete. If specified, all other arguments will be ignored
	// Required: false
	RuleID uint64 `url:"ruleId,omitempty"`

	// External start port number for the rule
	// Required: false
	PublicPortStart uint64 `url:"publicPortStart,omitempty"`

	// End port number (inclusive) for the ranged rule
	// Required: false
	PublicPortEnd uint64 `url:"publicPortEnd,omitempty"`

	// Internal base port number
	// Required: false
	LocalBasePort uint64 `url:"localBasePort,omitempty"`

	// Network protocol
	// Should be one of:
	//	- tcp
	//	- udp
	// Required: false
	Proto string `url:"proto,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq PFWDelRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// PFWDel delete port forward rule
func (c Compute) PFWDel(ctx context.Context, req PFWDelRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/pfwDel"

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
