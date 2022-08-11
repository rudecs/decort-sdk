package flipgroup

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type FlipGroup struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *FlipGroup {
	return &FlipGroup{
		client,
	}
}
