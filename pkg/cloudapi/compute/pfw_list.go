package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type PFWListRequest struct {
	ComputeID uint64 `url:"computeId"`
}

func (crq PFWListRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) PFWList(ctx context.Context, req PFWListRequest) (PFWList, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/pfwList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	pfwList := PFWList{}

	err = json.Unmarshal(res, &pfwList)
	if err != nil {
		return nil, err
	}

	return pfwList, nil
}
