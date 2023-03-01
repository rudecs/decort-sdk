package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update capacity limits
type UpdateCapacityLimitRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`
}

func (srq UpdateCapacityLimitRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// UpdateCapacityLimit updates SEP capacity limit
func (s SEP) UpdateCapacityLimit(ctx context.Context, req UpdateCapacityLimitRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/sep/updateCapacityLimit"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
