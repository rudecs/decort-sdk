package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get boot order
type BootOrderGetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`
}

func (crq BootOrderGetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// BootOrderGet gets actual compute boot order information
func (c Compute) BootOrderGet(ctx context.Context, req BootOrderGetRequest) ([]string, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/bootOrderGet"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	orders := make([]string, 0)

	err = json.Unmarshal(res, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
