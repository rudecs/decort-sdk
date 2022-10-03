package sizes

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Sizes struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Sizes {
	return &Sizes{
		client,
	}
}
