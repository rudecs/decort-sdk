package flipgroup

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about FLIPGroup
type GetRequest struct {
	// FLIPGroup ID
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId"`
}

func (frq GetRequest) validate() error {
	if frq.FLIPGroupID == 0 {
		return errors.New("field FLIPGroupID can not be empty or equal to 0")
	}

	return nil
}

// Get gets details of the specified Floating IP group
func (f FLIPGroup) Get(ctx context.Context, req GetRequest) (*ItemFLIPGroup, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudapi/flipgroup/get"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ItemFLIPGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
