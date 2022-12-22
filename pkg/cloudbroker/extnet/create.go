package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create external network
type CreateRequest struct {
	// External network name
	// Required: true
	Name string `url:"name"`

	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid"`

	// IP network CIDR
	// For example 192.168.0.0/24
	// Required: true
	IPCIDR string `url:"ipcidr"`

	// External network gateway IP address
	// Required: true
	Gateway string `url:"gateway"`

	// VLAN ID
	// Required: true
	VLANID uint64 `url:"vlanId"`

	// List of DNS addresses
	// Required: false
	DNS []string `url:"dns,omitempty"`

	// List of NTP addresses
	// Required: false
	NTP []string `url:"ntp,omitempty"`

	// IPs to check network availability
	// Required: false
	CheckIPs []string `url:"checkIps,omitempty"`

	// If true - platform DHCP server will not be created
	// Required: false
	Virtual bool `url:"virtual,omitempty"`

	// Optional description
	// Required: false
	Description string `url:"desc,omitempty"`

	// Start of IP range to be explicitly included
	// Required: false
	StartIP string `url:"startIP,omitempty"`

	// End of IP range to be explicitly included
	// Required: false
	EndIP string `url:"endIP,omitempty"`

	// IP to create VNFDev with
	// Required: false
	VNFDevIP string `url:"vnfdevIP,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint64 `url:"preReservationsNum,omitempty"`

	// OpenvSwith bridge name for ExtNet connection
	// Required: false
	OVSBridge string `url:"ovsBridge,omitempty"`
}

func (erq CreateRequest) validate() error {
	if erq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if erq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if erq.IPCIDR == "" {
		return errors.New("validation-error: field IPCIDR must be set")
	}
	if erq.Gateway == "" {
		return errors.New("validation-error: field Gateway must be set")
	}
	if erq.VLANID == 0 {
		return errors.New("validation-error: field VLANID must be set")
	}

	return nil
}

// Create creates new external network into platform
func (e ExtNet) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/extnet/create"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
