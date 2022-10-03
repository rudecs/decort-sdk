package extnet

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type ExtNet struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *ExtNet {
	return &ExtNet{
		client,
	}
}
