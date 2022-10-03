package bservice

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type GroupResizeRequest struct {
	ServiceID   uint64 `url:"serviceId"`
	CompGroupID uint64 `url:"compgroupId"`
	Count       int64  `url:"count"`
	Mode        string `url:"mode"`
}

func (bsrq GroupResizeRequest) Validate() error {
	if bsrq.ServiceID == 0 {
		return errors.New("field ServiceID can not be empty or equal to 0")
	}

	if bsrq.CompGroupID == 0 {
		return errors.New("field CompGroupID can not be empty or equal to 0")
	}

	if bsrq.Mode == "RELATIVE" && bsrq.Count == 0 {
		return errors.New("field Count can not be equal to 0 if Mode if 'RELATIVE'")
	}

	if !validators.StringInSlice(bsrq.Mode, []string{"RELATIVE", "ABSOLUTE"}) {
		return errors.New("field Mode can only be one of 'RELATIVE' or 'ABSOLUTE'")
	}

	return nil
}

func (b BService) GroupResize(ctx context.Context, req GroupResizeRequest) (uint64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/groupResize"
	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(string(res), 10, 64)
}
