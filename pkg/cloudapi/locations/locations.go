package locations

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Locations struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Locations {
	return &Locations{
		client,
	}
}
