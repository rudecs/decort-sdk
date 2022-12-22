package sep

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get SEP pool config by name
type GetPoolRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Pool name
	// Required: true
	PoolName string `url:"pool_name"`
}

func (srq GetPoolRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if srq.PoolName == "" {
		return errors.New("validation-error: PoolName must be set")
	}

	return nil
}

// GetPool gets SEP pool config by name
func (s SEP) GetPool(ctx context.Context, req GetPoolRequest) (*RecordPool, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/sep/getPool"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordPool{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
