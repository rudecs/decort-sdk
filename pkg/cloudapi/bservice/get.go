package bservice

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get detailed information about service
type GetRequest struct {
	// ID of the service to query information
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId"`
}

func (bsrq GetRequest) validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	return nil
}

// Get gets detailed specifications for the BasicService.
func (b BService) Get(ctx context.Context, req GetRequest) (*RecordBasicService, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/bservice/get"

	bsRaw, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordBasicService{}

	err = json.Unmarshal(bsRaw, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
