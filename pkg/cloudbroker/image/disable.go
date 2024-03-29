package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for disable image
type DisableRequest struct {
	// ID of image to be disabled
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`
}

func (irq DisableRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// Disable disables image
func (i Image) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/disable"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
