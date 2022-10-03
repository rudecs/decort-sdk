package disks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListUnattachedRequest struct {
	AccountID uint64 `url:"accountId"`
}

func (drq ListUnattachedRequest) Validate() error {
	if drq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}

	return nil
}

func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest) (ListUnattached, error) {
	url := "/cloudbroker/disks/listUnattached"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	diskList := ListUnattached{}

	err = json.Unmarshal(res, &diskList)
	if err != nil {
		return nil, err
	}

	return diskList, nil

}
