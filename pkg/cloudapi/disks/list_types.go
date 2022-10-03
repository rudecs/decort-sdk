package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListTypesRequest struct {
	Detailed bool `url:"detailed"`
}

func (d Disks) ListTypes(ctx context.Context, req ListTypesRequest) ([]interface{}, error) {
	url := "/cloudapi/disks/listTypes"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	typesList := make([]interface{}, 0)

	err = json.Unmarshal(res, &typesList)
	if err != nil {
		return nil, err
	}

	return typesList, nil

}
