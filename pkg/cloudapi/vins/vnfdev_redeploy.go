package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for redeploy VNFDevs
type VNFDevRedeployRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId"`
}

func (vrq VNFDevRedeployRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

// VNFDevRedeploy redeploy VINS VNFDevs
func (v VINS) VNFDevRedeploy(ctx context.Context, req VNFDevRedeployRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/vins/vnfdevRedeploy"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
