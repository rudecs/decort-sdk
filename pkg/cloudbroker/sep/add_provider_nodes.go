package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for add provider nodes
type AddProviderNodesRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id"`

	// List of node IDs
	// Required: true
	ProviderNIDs []uint64 `url:"provider_nids" json:"provider_nids"`
}

func (srq AddProviderNodesRequest) validate() error {
	if srq.SEPID == 0 {
		return errors.New("validation-error: field SEPID must be set")
	}
	if len(srq.ProviderNIDs) == 0 {
		return errors.New("validation-error: field ProviderNIDs must be set")
	}

	return nil
}

// AddProviderNodes add provider nodes to SEP parameters
func (s SEP) AddProviderNodes(ctx context.Context, req AddProviderNodesRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/sep/addProviderNodes"

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
