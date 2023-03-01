package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for reset config
type ConfigResetRequest struct {
	// ID of the load balancer instance to ConfigReset
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`
}

func (lbrq ConfigResetRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	return nil
}

// ConfigReset reset current software configuration of the specified load balancer.
// Warning: this action cannot be undone!
func (l LB) ConfigReset(ctx context.Context, req ConfigResetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/configReset"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
