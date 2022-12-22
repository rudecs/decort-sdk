package vins

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for IP reserve
type IPReserveRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// Type of the reservation
	// Should be one of:
	//	- DHCP
	//	- VIP
	//	- EXCLUDE
	// Required: true
	Type string `url:"type"`

	// IP address to use. Non-empty string is required for type "EXCLUDE".
	// Ignored for types "DHCP" and "VIP".
	// Required: false
	IPAddr string `url:"ipAddr,omitempty"`

	// MAC address to associate with IP reservation.
	// Ignored for type "EXCLUDE",
	// non-empty string is required for "DHCP" and "VIP"
	// Required: false
	MAC string `url:"mac,omitempty"`

	// ID of the compute, associated with this reservation of type "DHCP".
	// Ignored for other types
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (vrq IPReserveRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}
	if vrq.Type == "" {
		return errors.New("validation-error: field Type must be set")
	}
	validate := validators.StringInSlice(vrq.Type, []string{"DHCP", "VIP", "EXCLUDED"})
	if !validate {
		return errors.New("'type' should be 'DHCP', 'VIP' or 'EXCLUDED'")
	}

	return nil
}

// IPReserve creates reservation on ViNS DHCP
func (v VINS) IPReserve(ctx context.Context, req IPReserveRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/vins/ipReserve"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
