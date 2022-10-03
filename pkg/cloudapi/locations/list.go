package locations

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	Page uint64 `url:"page"`
	Size uint64 `url:"size"`
}

func (l Locations) List(ctx context.Context, req ListRequest) (LocationsList, error) {
	url := "/locations/list"
	prefix := "/cloudapi"

	url = prefix + url
	locationsListRaw, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	locationsList := LocationsList{}
	err = json.Unmarshal(locationsListRaw, &locationsList)
	if err != nil {
		return nil, err
	}

	return locationsList, nil

}
