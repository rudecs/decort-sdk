package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for unset compute CI
type ComputeCIUnsetRequest struct {
	// ID of the image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`
}

func (irq ComputeCIUnsetRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// ComputeCIUnset unset compute CI ID from image
func (i Image) ComputeCIUnset(ctx context.Context, req ComputeCIUnsetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/сomputeciUnset"

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
