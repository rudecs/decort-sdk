package kvmx86

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type KVMX86 struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *KVMX86 {
	return &KVMX86{
		client,
	}
}
