package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteImagesRequest struct {
	ImageIDs    []uint64 `url:"imageIds"`
	Reason      string   `url:"reason"`
	Permanently bool     `url:"permanently"`
}

func (irq DeleteImagesRequest) Validate() error {
	if len(irq.ImageIDs) == 0 {
		return errors.New("validation-error: field ImageIDs must be set")
	}
	if irq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

func (i Image) DeleteImages(ctx context.Context, req DeleteImagesRequest) (bool, error) {
	err := req.Validate()
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
