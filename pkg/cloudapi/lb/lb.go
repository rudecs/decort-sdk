package lb

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type LB struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *LB {
	return &LB{
		client,
	}
}
