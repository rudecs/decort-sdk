package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for set boot order
type BootOrderSetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// List of boot devices
	// Should be one of:
	//	- cdrom
	//	- network
	//	- hd
	// Required: true
	Order []string `url:"order" json:"order"`
}

func (crq BootOrderSetRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if len(crq.Order) == 0 {
		return errors.New("validation-error: field Order must be set")
	}
	for _, value := range crq.Order {
		if validate := validators.StringInSlice(value, []string{"cdrom", "network", "hd"}); !validate {
			return errors.New("validation-error: field ImageType can be cdrom, network, hd")
		}
	}

	return nil
}

// BootOrderSet sets compute boot order
func (c Compute) BootOrderSet(ctx context.Context, req BootOrderSetRequest) ([]string, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/compute/bootOrderSet"

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
