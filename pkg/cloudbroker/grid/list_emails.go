package grid

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListEmails returns list of email addresses of users
func (g Grid) ListEmails(ctx context.Context) ([]string, error) {
	url := "/cloudbroker/grid/listEmails"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
