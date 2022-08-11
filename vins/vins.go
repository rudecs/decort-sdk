package vins

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Vins struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Vins {
	return &Vins{
		client,
	}
}
