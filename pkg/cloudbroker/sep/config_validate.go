package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for validate config
type ConfigValidateRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`

	// Storage provider config
	// Required: true
	Config string `url:"config" json:"config"`
}

func (srq ConfigValidateRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.Config == "" {
		return errors.New("validation-error: feold Config must be set")
	}

	return nil
}

// ConfigValidate verify config for the SEP
func (s SEP) ConfigValidate(ctx context.Context, req ConfigValidateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/configValidate"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
