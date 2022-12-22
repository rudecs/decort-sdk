package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update QOS
type DefaultQOSUpdateRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`

	// Internal traffic, kbit
	// Required: false
	IngressRate uint64 `url:"ingress_rate,omitempty"`

	// Internal traffic burst, kbit
	// Required: false
	IngressBurst uint64 `url:"ingress_burst,omitempty"`

	// External traffic rate, kbit
	// Required: false
	EgressRate uint64 `url:"egress_rate,omitempty"`
}

func (erq DefaultQOSUpdateRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DefaultQOSUpdate updates default qos values
func (e ExtNet) DefaultQOSUpdate(ctx context.Context, req DefaultQOSUpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/defaultQosUpdate"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
