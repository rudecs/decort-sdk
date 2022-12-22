package sep

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get consumption info
type ConsumptionRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`
}

func (srq ConsumptionRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// Consumption get SEP consumption info
func (s SEP) Consumption(ctx context.Context, req ConsumptionRequest) (*RecordConsumption, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/sep/consumption"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordConsumption{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
