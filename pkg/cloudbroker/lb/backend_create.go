package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

// Request struct for create backend
type BackendCreateRequest struct {
	// ID of the load balancer instance to backendCreate
	// Required: true
	LBID uint64 `url:"lbId"`

	// Must be unique among all backends of this load balancer - name of the new backend to create
	// Required: true
	BackendName string `url:"backendName"`

	// Algorithm
	// Should be one of:
	//	- roundrobin
	//	- static-rr
	//	- leastconn
	// Required: false
	Algorithm string `url:"algorithm,omitempty"`

	// Interval in milliseconds between two consecutive availability
	// checks of the server that is considered available
	// Required: false
	Inter uint64 `url:"inter,omitempty"`

	// Interval in milliseconds between two consecutive checks to
	// restore the availability of a server that is currently considered unavailable
	// Required: false
	DownInter uint64 `url:"downinter,omitempty"`

	// Number of checks that the server must pass in order to get the available status
	// and be included in the balancing scheme again
	// Required: false
	Rise uint64 `url:"rise,omitempty"`

	// Number of consecutive failed availability checks,
	// after which the previously considered available server receives the status of
	// unavailable and is temporarily excluded from the balancing scheme
	// Required: false
	Fall uint64 `url:"fall,omitempty"`

	// Interval in milliseconds from the moment the server receives the available status,
	// after which the number of actually allowed connections to this server will be returned to 100% of the set limit
	// Required: false
	SlowStart uint64 `url:"slowstart,omitempty"`

	// Limit of simultaneous connections to the server. When this limit is reached,
	// the server is temporarily excluded from the balancing scheme
	// Required: false
	MaxConn uint64 `url:"maxconn,omitempty"`

	// Limit of connections waiting in the queue.
	// When this limit is reached, all subsequent connections will be forwarded to other servers
	// Required: false
	MaxQueue uint64 `url:"maxqueue,omitempty"`

	// Server weight for use in weight balancing algorithms
	// Required: false
	Weight uint64 `url:"weight,omitempty"`
}

func (lbrq BackendCreateRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID must be set")
	}
	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName must be set")
	}
	validate := validators.StringInSlice(lbrq.Algorithm, []string{"roundrobin", "static-rr", "leastconn"})
	if !validate {
		return errors.New("validation-error: field Algorithm must be one of roundrobin, static-rr, leastconn")
	}

	return nil
}

// BackendCreate creates new backend on the specified load balancer
func (lb LB) BackendCreate(ctx context.Context, req BackendCreateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/lb/backendCreate"

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
