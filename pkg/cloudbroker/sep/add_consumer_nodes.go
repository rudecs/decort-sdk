package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add consumer nodes
type AddConsumerNodesRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// List of nodes IDs
	// Required: true
	ConsumerNIDs []uint64 `url:"consumer_nids"`
}

func (srq AddConsumerNodesRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if len(srq.ConsumerNIDs) == 0 {
		return errors.New("validation-error: field ConsumerNIDs must be set")
	}

	return nil
}

// AddConsumerNodes add consumer nodes to SEP parameters
func (s SEP) AddConsumerNodes(ctx context.Context, req AddConsumerNodesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/addConsumerNodes"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
