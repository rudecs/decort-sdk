package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for rename image
type RenameRequest struct {
	// ID of the virtual image to rename
	// Required: true
	ImageID uint64 `url:"imageId"`

	// New name
	// Required: true
	Name string `url:"name"`
}

func (irq RenameRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}

	return nil
}

// Rename renames image by ID
func (i Image) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/rename"

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
