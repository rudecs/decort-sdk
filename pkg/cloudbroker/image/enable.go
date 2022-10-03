package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type EnableRequest struct {
	ImageID uint64 `url:"imageId"`
}

func (irq EnableRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must br set")
	}

	return nil
}

func (i Image) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/enable"

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