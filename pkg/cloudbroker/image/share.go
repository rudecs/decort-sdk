package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for share image
type ShareRequest struct {
	// ID of the image to share
	// Required: true
	ImageId uint64 `url:"imageId" json:"imageId"`

	// List of account IDs
	// Required: true
	AccountIDs []uint64 `url:"accounts" json:"accounts"`
}

func (irq ShareRequest) validate() error {
	if irq.ImageId == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if len(irq.AccountIDs) == 0 || irq.AccountIDs == nil {
		return errors.New("validation-error: field must be set")
	}

	return nil
}

// Share shares image with accounts
func (i Image) Share(ctx context.Context, req ShareRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/share"

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
