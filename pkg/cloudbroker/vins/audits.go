package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get audits
type AuditsRequest struct {
	// ID of the VINS
	// Required: true
	VINSID uint64 `url:"vinsId"`
}

func (vrq AuditsRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// Audits gets audit records for the specified VINS object
func (v VINS) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/vins/audits"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
