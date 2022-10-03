package vins

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ExtNetListRequest struct {
	VINSID uint64 `url:"vinsId"`
}

func (vrq ExtNetListRequest) Validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID can not be empty or equal to 0")
	}

	return nil
}

func (v VINS) ExtNetList(ctx context.Context, req ExtNetListRequest) (ExtNetList, error) {
	url := "/cloudapi/vins/extNetList"

	extnetListRaw, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	extnetList := ExtNetList{}
	err = json.Unmarshal(extnetListRaw, &extnetList)
	if err != nil {
		return nil, err
	}

	return extnetList, nil

}
