package grid

import (
	"context"
	"errors"
	"net/http"
)

// Request struct for get platform snapshot with additional diagnosis
type GetDiagnosisRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid"`
}

func (grq GetDiagnosisRequest) validate() error {
	if grq.GID == 0 {
		return errors.New("validation-error: field GID must be set")
	}

	return nil
}

// GetDiagnosis get platform snapshot with additional diagnosis info like a logs, etc
func (g Grid) GetDiagnosis(ctx context.Context, req GetDiagnosisRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/grid/getDiagnosis"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// GetDiagnosisGET get platform snapshot with additional diagnosis info like a logs, etc
func (g Grid) GetDiagnosisGET(ctx context.Context, req GetDiagnosisRequest) (string, error) {
	err := req.validate()
	if err != nil {
		return "", err
	}

	url := "/cloudbroker/grid/getDiagnosis"

	res, err := g.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
