package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for delete server definition
type BackendServerDeleteRequest struct {
	// ID of the load balancer instance to BackendServerDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId"`

	// Must match one of the existing backens - name of the backend to add servers to
	// Required: true
	BackendName string `url:"backendName" json:"backendName"`

	// Must be unique among all servers defined for this backend - name of the server definition to add
	// Required: true
	ServerName string `url:"serverName" json:"serverName"`
}

func (lbrq BackendServerDeleteRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}

	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName must be set")
	}

	if lbrq.ServerName == "" {
		return errors.New("validation-error: field ServerName must be set")
	}

	return nil
}

// BackendServerDelete deletes server definition from the backend on the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) BackendServerDelete(ctx context.Context, req BackendServerDeleteRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/backendServerDelete"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
