package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type LinkRequest struct {
	ImageID  uint64 `url:"imageId"`
	TargetID uint64 `url:"targetId"`
}

func (irq LinkRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.TargetID == 0 {
		return errors.New("validation-error: field TargetID must be set")
	}

	return nil
}

func (i Image) Link(ctx context.Context, req LinkRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/link"

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