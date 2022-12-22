package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for insert config
type ConfigInsertRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Storage provider config
	// Required: true
	Config string `url:"config"`
}

func (srq ConfigInsertRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.Config == "" {
		return errors.New("validation-error: feold Config must be set")
	}

	return nil
}

// ConfigInsert insert config to SEP
func (s SEP) ConfigInsert(ctx context.Context, req ConfigInsertRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/configInsert"

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
