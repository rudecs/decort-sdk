package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get list VINS external network connections
type ExtNetListRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`
}

func (vrq ExtNetListRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// ExtNetList show list of VINS external network connections
func (v VINS) ExtNetList(ctx context.Context, req ExtNetListRequest) (ListExtNets, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/vins/extNetList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListExtNets{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil

}
