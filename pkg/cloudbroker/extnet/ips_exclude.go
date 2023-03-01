package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for exclude list IPs
type IPsExcludeRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`

	// List of IPs for exclude from external network
	// Required: true
	IPs []string `url:"ips" json:"ips"`
}

func (erq IPsExcludeRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}
	if len(erq.IPs) == 0 {
		return errors.New("validation-error: field IPs must be set")
	}

	return nil
}

// IPsExclude exclude list IPs from external network pool
func (e ExtNet) IPsExclude(ctx context.Context, req IPsExcludeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/ipsExclude"

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
