package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ComputeCISetRequest struct {
	ImageID     uint64 `url:"imageId"`
	ComputeCIID uint64 `url:"computeciId"`
}

func (irq ComputeCISetRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.ComputeCIID == 0 {
		return errors.New("validation-error: field ComputeCIID must be set")
	}

	return nil
}

func (i Image) ComputeCISet(ctx context.Context, req ComputeCISetRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/computeciSet"

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
