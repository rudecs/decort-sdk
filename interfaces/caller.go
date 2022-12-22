package interfaces

import "context"

// Interface for sending requests to platform
type Caller interface {
	// DecortApiCall method for sending requests to the platform
	DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error)
}
