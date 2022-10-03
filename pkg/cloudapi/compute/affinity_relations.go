package compute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AffinityRelationsRequest struct {
	ComputeID     uint64 `url:"computeId"`
	AffinityLabel string `url:"affinityLabel"`
}

func (crq AffinityRelationsRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	return nil
}

func (c Compute) AffinityRelations(ctx context.Context, req AffinityRelationsRequest) (*AffinityRelations, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/compute/affinityRelations"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	relations := &AffinityRelations{}

	err = json.Unmarshal(res, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}
