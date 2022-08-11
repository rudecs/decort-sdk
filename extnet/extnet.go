package extnet

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Extnet struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Extnet {
	return &Extnet{
		client,
	}
}
