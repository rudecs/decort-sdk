package interfaces

import "context"

type Caller interface {
	DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error)
}
