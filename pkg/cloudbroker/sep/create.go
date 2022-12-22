package sep

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create SEP object
type CreateRequest struct {
	// Grid ID
	// Required: true
	GID uint64 `url:"gid"`

	// SEP name
	// Required: true
	Name string `url:"name"`

	// Type of storage
	// Required: true
	SEPType string `url:"sep_type"`

	// Description
	// Required: false
	Description string `url:"description,omitempty"`

	// SEP config
	// Required: false
	Config string `url:"config,omitempty"`

	// List of provider node IDs
	// Required: false
	ProviderNIDs []uint64 `url:"provider_nids,omitempty"`

	// List of consumer node IDs
	// Required: false
	ConsumerNIDs []uint64 `url:"consumer_nids,omitempty"`

	// Enable SEP after creation
	// Required: false
	Enable bool `url:"enable,omitempty"`
}

func (srq CreateRequest) validate() error {
	if srq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}
	if srq.Name == "" {
		return errors.New("validation-error: field Name must be set")
	}
	if srq.SEPType == "" {
		return errors.New("validation-error: field SEPType must be set")
	}

	return nil
}

// Create creates SEP object
func (s SEP) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/sep/create"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
