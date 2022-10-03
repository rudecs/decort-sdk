package lb

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type BackendCreateRequest struct {
	LBID        uint64 `url:"lbId"`
	BackendName string `url:"backendName"`
	Algorithm   string `url:"algorithm,omitempty"`
	Inter       uint64 `url:"inter,omitempty"`
	DownInter   uint64 `url:"downinter,omitempty"`
	Rise        uint   `url:"rise,omitempty"`
	Fall        uint   `url:"fall,omitempty"`
	SlowStart   uint64 `url:"slowstart,omitempty"`
	MaxConn     uint   `url:"maxconn,omitempty"`
	MaxQueue    uint   `url:"maxqueue,omitempty"`
	Weight      uint   `url:"weight,omitempty"`
}

func (lbrq BackendCreateRequest) Validate() error {
	if lbrq.LBID == 0 {
		return errors.New("validation-error: field LBID can not be empty or equal to 0")
	}

	if lbrq.BackendName == "" {
		return errors.New("validation-error: field BackendName can not be empty")
	}

	return nil
}

func (l LB) BackendCreate(ctx context.Context, req BackendCreateRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/lb/backendCreate"

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
