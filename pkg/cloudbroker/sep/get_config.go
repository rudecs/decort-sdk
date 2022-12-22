package sep

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get SEP config
type GetConfigRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`
}

func (srq GetConfigRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// GetConfig gets SEP config
func (s SEP) GetConfig(ctx context.Context, req GetConfigRequest) (*SEPConfig, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/sep/getConfig"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := SEPConfig{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
