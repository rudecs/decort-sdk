package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete CD-ROM image
type DeleteCDROMImageRequest struct {
	// ID of the CD-ROM image to delete
	// Required: true
	ImageID uint64 `url:"imageId"`

	// Whether to completely delete the CD-ROM image, needs to be unused
	// Required: true
	Permanently bool `url:"permanently"`
}

func (irq DeleteCDROMImageRequest) validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

// DeleteCDROMImage delete a CD-ROM image
func (i Image) DeleteCDROMImage(ctx context.Context, req DeleteCDROMImageRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/deleteCDROMImage"

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
