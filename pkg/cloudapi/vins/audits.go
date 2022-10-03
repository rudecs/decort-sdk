package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AuditsRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq AuditsRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) Audits(ctx context.Context, req AuditsRequest) (VINSAuditsList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/vins/audits"

	auditsRaw, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	audits := VINSAuditsList{}

	err = json.Unmarshal(auditsRaw, &audits)
	if err != nil {
		return nil, err
	}

	return audits, nil
}
