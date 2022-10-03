package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	ImageID     uint64 `url:"imageId"`
	Reason      string `url:"reason"`
	Permanently bool   `url:"permanently,omitempty"`
}

func (irq DeleteRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}
	if irq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}

	return nil
}

func (i Image) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/delete"

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
