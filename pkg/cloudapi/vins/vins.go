package vins

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type VINS struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *VINS {
	return &VINS{
		client,
	}
}
