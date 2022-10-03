package k8s

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type K8S struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *K8S {
	return &K8S{
		client,
	}
}
