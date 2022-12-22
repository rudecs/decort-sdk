package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for set compute CI
type ComputeCISetRequest struct {
	// ID of the image
	// Required: true
	ImageID uint64 `url:"imageId"`

	// ID of the compute CI
	// Required: true
	ComputeCIID uint64 `url:"computeciId"`
}

func (irq ComputeCISetRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.ComputeCIID == 0 {
		return errors.New("validation-error: field ComputeCIID must be set")
	}

	return nil
}

// ComputeCISet set compute CI ID for image
func (i Image) ComputeCISet(ctx context.Context, req ComputeCISetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/computeciSet"

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
