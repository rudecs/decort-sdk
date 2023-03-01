package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about VINS
type GetRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (vrq GetRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// Get gets information about VINS by ID
func (v VINS) Get(ctx context.Context, req GetRequest) (*RecordVINS, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/vins/get"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordVINS{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
