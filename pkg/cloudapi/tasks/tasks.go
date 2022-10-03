package tasks

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Tasks struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Tasks {
	return &Tasks{
		client,
	}
}
