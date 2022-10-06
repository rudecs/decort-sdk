package rg

import "github.com/rudecs/decort-sdk/interfaces"

type RG struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *RG {
	return &RG{
		client: client,
	}
}
