package bservice

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GroupGetRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	CompGroupID uint64 `url:"compgroupId"`
}

func (bsrq GroupGetRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) GroupGet(ctx context.Context, req GroupGetRequest) (*Group, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	url := "/cloudapi/bservice/groupGet"
	groupRaw, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	group := &Group{}
	if err := json.Unmarshal(groupRaw, group); err != nil {
		return nil, err
	}

	return group, nil
}
