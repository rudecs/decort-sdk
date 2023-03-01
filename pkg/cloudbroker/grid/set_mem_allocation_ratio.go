package grid

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set memory allocation
type SetMemAllocationRatioRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gridId" json:"gridId"`

	// Allocation ratio
	// Required: true
	Ratio float64 `url:"ratio" json:"ratio"`
}

func (grq SetMemAllocationRatioRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if grq.Ratio == 0.0 {
		return errors.New("validation-error: field Ratio must be set")
	}

	return nil
}

// SetMemAllocationRatio sets memory allocation ratio
func (g Grid) SetMemAllocationRatio(ctx context.Context, req SetMemAllocationRatioRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/grid/setMemAllocationRatio"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
