package kvmppc

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type KVMPPC struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *KVMPPC {
	return &KVMPPC{
		client,
	}
}
