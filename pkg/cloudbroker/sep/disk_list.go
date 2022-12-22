package sep

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list of disk IDs
type DiskListRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// Pool name
	// Required: false
	PoolName string `url:"pool_name,omitempty"`
}

func (srq DiskListRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// DiskList get list of disk IDs, who use this SEP and pool (if provided)
func (s SEP) DiskList(ctx context.Context, req DiskListRequest) ([]uint64, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/sep/diskList"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]uint64, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
