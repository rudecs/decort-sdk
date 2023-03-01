package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get image details
type GetRequest struct {
	// ID of image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`
}

func (irq GetRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// Get get image details by ID
func (i Image) Get(ctx context.Context, req GetRequest) (*RecordImage, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/image/get"

	info := RecordImage{}

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
