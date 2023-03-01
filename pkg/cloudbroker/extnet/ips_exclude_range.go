package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for exclude range of IPs
type IPsExcludeRangeRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`

	// Starting IP
	// Required: true
	IPStart string `url:"ip_start" json:"ip_start"`

	// Ending IP
	// Required: true
	IPEnd string `url:"ip_end" json:"ip_end"`
}

func (erq IPsExcludeRangeRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}
	if erq.IPStart == "" {
		return errors.New("validation-error: field IPStart must be set")
	}
	if erq.IPEnd == "" {
		return errors.New("validation-error: field IPEnd must be set")
	}

	return nil
}

// IPsExcludeRange exclude range of IPs from external network pool
func (e ExtNet) IPsExcludeRange(ctx context.Context, req IPsExcludeRangeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/ipsExcludeRange"

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
