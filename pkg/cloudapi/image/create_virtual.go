package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type CreateVirtualRequest struct {
	Name     string `url:"name"`
	TargetID uint64 `url:"targetId"`
}

func (irq CreateVirtualRequest) Validate() error {
	if irq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}

	if irq.TargetID == 0 {
		return errors.New("validation-error: field TargetID can not be empty or equal to 0")
	}

	return nil
}

func (i Image) CreateVirtual(ctx context.Context, req CreateVirtualRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/image/createVirtual"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil

}
