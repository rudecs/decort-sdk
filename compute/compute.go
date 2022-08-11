package compute

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Compute struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Compute {
	return &Compute{
		client,
	}
}
