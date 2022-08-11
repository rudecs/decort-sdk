package k8ci

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type K8CI struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *K8CI {
	return &K8CI{
		client,
	}
}
