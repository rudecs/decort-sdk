package image

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type UpdateNodesRequest struct {
	ImageID       uint64   `url:"imageId"`
	EnabledStacks []uint64 `url:"enabledStacks,omitempty"`
}

func (irq UpdateNodesRequest) Validate() error {
	if irq.ImageID == 0 {
		return errors.New("validation-error: field ImageID must be set")
	}

	return nil
}

func (i Image) UpdateNodes(ctx context.Context, req UpdateNodesRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/image/updateNodes"

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
