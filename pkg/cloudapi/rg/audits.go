package rg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get audit
type AuditsRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId"`
}

func (rgrq AuditsRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

// Audits gets audit records for the specified resource group object
func (r RG) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/rg/audits"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
