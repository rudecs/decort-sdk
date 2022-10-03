package cloudapi

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type CloudApi struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *CloudApi {
	return &CloudApi{
		client: client,
	}
}
