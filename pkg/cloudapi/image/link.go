package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for link virtual image to another image
type LinkRequest struct {
	// ID of the virtual image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// ID of real image to link this virtual image to
	// Required: true
	TargetID uint64 `url:"targetId" json:"targetId"`
}

func (irq LinkRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}
	if irq.TargetID == 0 {
		return errors.New("validation-error: field TargetID can not be empty or equal to 0")
	}

	return nil
}

// Link links virtual image to another image in the platform
func (i Image) Link(ctx context.Context, req LinkRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/image/link"

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
