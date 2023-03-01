package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get detailed information about image
type GetRequest struct {
	// ID of image to get
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// If set to False returns only images in status CREATED
	// Required: false
	ShowAll bool `url:"show_all,omitempty" json:"show_all,omitempty"`
}

func (irq GetRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}

	return nil
}

// Get gets image by ID.
// Returns image if user has rights on it
func (i Image) Get(ctx context.Context, req GetRequest) (*RecordImage, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/image/get"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordImage{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
