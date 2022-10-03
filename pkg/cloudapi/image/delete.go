package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	ImageID     uint64 `url:"imageId"`
	Permanently bool   `url:"permanently"`
}

func (irq DeleteRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}

	return nil
}

func (i Image) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/image/delete"

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
