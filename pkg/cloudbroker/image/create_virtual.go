package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create virtual image
type CreateVirtualRequest struct {
	// Name of the virtual image to create
	// Required: true
	Name string `url:"name" json:"name"`

	// ID of real image to link this virtual image to upon creation
	// Required: true
	TargetID uint64 `url:"targetId" json:"targetId"`
}

func (irq CreateVirtualRequest) validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if irq.TargetID == 0 {
		return errors.New("validation-error: field TargetID must be set")
	}

	return nil
}

// CreateVirtual creates virtual image
func (i Image) CreateVirtual(ctx context.Context, req CreateVirtualRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/image/createVirtual"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
