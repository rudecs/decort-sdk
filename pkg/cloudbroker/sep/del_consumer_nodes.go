package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for exclude consumer nodes
type DelConsumerNodesRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id"`

	// List of consumer node IDs
	// Required: true
	ConsumerNIDs []uint64 `url:"consumer_nids"`
}

func (srq DelConsumerNodesRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if len(srq.ConsumerNIDs) == 0 {
		return errors.New("validation-error: field ConsumerNIDs must be set")
	}

	return nil
}

// DelConsumerNodes exclude consumer nodes from SEP parameters
func (s SEP) DelConsumerNodes(ctx context.Context, req DelConsumerNodesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/delConsumerNodes"

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
