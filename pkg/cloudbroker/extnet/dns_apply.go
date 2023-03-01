package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set new DNS
type DNSApplyRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`

	// List of DNS to apply
	// Required: false
	DNSList []string `url:"dns_list,omitempty" json:"dns_list,omitempty"`
}

func (erq DNSApplyRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// DNSApply sets new DNS
func (e ExtNet) DNSApply(ctx context.Context, req DNSApplyRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/dnsApply"

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
