package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for decommission
type DecommissionRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`

	// Clear disks and images physically
	// Required: false
	ClearPhisically bool `url:"clear_physically,omitempty" json:"clear_physically,omitempty"`
}

func (srq DecommissionRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// Decommission unlink everything that exists from SEP
func (s SEP) Decommission(ctx context.Context, req DecommissionRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/decommission"

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
