package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteCDROMImageRequest struct {
	ImageID     uint64 `url:"imageId"`
	Permanently bool   `url:"permanently"`
}

func (irq DeleteCDROMImageRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (i Image) DeleteCDROMImage(ctx context.Context, req DeleteCDROMImageRequest) (bool, error) {
	err := req.Validate()
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
