package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete image
type DeleteRequest struct {
	// ID of the image to delete
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Reason for action
	// Required: true
	Reason string `url:"reason" json:"reason"`

	// Whether to completely delete the image
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

func (irq DeleteRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

// Delete deletes image by ID
func (i Image) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/delete"

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
