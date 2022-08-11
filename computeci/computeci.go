package computeci

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type ComputeCI struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *ComputeCI {
	return &ComputeCI{
		client,
	}
}
