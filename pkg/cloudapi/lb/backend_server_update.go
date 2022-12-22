package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for update server
type BackendServerUpdateRequest struct {
	// ID of the load balancer instance to BackendServerAdd
	// Required: true
	LBID uint64 `url:"lbId"`

	// Must match one of the existing backens - name of the backend to add servers to
	// Required: true
	BackendName string `url:"backendName"`

	// Must be unique among all servers defined for this backend - name of the server definition to add
	// Required: true
	ServerName string `url:"serverName"`

	// IP address of the server
	// Required: true
	Address string `url:"address"`

	// Port number on the server
	// Required: true
	Port uint64 `url:"port"`

	// Set to disabled if this server should be used regardless of its state
	// Required: false
	Check string `url:"check,omitempty"`

	// Interval in milliseconds between two consecutive availability checks of the server that is considered available
	// Required: false
	Inter uint64 `url:"inter,omitempty"`

	// Interval in milliseconds between two consecutive checks to restore
	// the availability of a server that is currently considered unavailable
	// Required: false
	DownInter uint64 `url:"downinter,omitempty"`

	// Number of checks that the server must pass in order to get
	// the available status and be included in the balancing scheme again
	// Required: false
	Rise uint64 `url:"rise,omitempty"`

	// Number of consecutive failed availability checks,
	// after which the previously considered available server receives the status of unavailable and
	// is temporarily excluded from the balancing scheme
	// Required: false
	Fall uint64 `url:"fall,omitempty"`

	// Interval in milliseconds from the moment the server receives the available status,
	// after which the number of actually allowed connections to this server will be returned to 100% of the set limit
	// Required: false
	SlowStart uint64 `url:"slowstart,omitempty"`

	// Limit of simultaneous connections to the server. When this limit is reached, the server is temporarily excluded from the balancing scheme
	// Required: false
	MaxConn uint64 `url:"maxconn,omitempty"`

	// Limit of connections waiting in the queue. When this limit is reached, all subsequent connections will be forwarded to other servers
	// Required: false
	MaxQueue uint64 `url:"maxqueue,omitempty"`

	// Server weight for use in weight balancing algorithms
	// Required: false
	Weight uint64 `url:"weight,omitempty"`
}

func (lbrq BackendServerUpdateRequest) validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}
	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}
	if lbrq.ServerName == "" {
		return errors.New("validation-error: field ServerName can not be empty")
	}
	if lbrq.Address == "" {
		return errors.New("validation-error: field Address can not be empty")
	}
	if lbrq.Port == 0 {
		return errors.New("validation-error: field Port can not be empty or equal to 0")
	}

	return nil
}

// BackendServerUpdate updates server definition on the backend of load balancer
func (l LB) BackendServerUpdate(ctx context.Context, req BackendServerUpdateRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/backendServerUpdate"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
