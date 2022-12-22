package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list port forwards
type PFWListRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq PFWListRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

// PFWList gets compute port forwards list
func (c Compute) PFWList(ctx context.Context, req PFWListRequest) (ListPFWs, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/pfwList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListPFWs{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
