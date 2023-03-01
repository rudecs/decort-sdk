package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete images
type DeleteImagesRequest struct {
	// List of images to be deleted
	// Required: true
	ImageIDs []uint64 `url:"imageIds" json:"imageIds"`

	// Reason for action
	// Required: true
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// Whether to completely delete the images
	// Required: true
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

func (irq DeleteImagesRequest) validate() error {
	if len(irq.ImageIDs) == 0 {
		return errors.New("validation-error: field ImageIDs must be set")
	}

	return nil
}

// DeleteImages deletes images
func (i Image) DeleteImages(ctx context.Context, req DeleteImagesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/deleteImages"

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
