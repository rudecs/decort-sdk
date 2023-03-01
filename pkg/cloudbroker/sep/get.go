package sep

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get SEP parameters
type GetRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`
}

func (srq GetRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}

	return nil
}

// Get gets SEP parameters
func (s SEP) Get(ctx context.Context, req GetRequest) (*RecordSEP, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/sep/get"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordSEP{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
