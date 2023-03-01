package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set new NTP
type NTPApplyRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id"`

	// List of NTP to apply
	// Required: false
	NTPList []string `url:"ntp_list,omitempty" json:"ntp_list,omitempty"`
}

func (erq NTPApplyRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}

	return nil
}

// NTPApply sets new NTP
func (e ExtNet) NTPApply(ctx context.Context, req NTPApplyRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/ntpApply"

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
