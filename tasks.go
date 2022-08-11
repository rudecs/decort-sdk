package client

import (
	"github.com/rudecs/decort-sdk/tasks"
)

func (dc *decortClient) Tasks() *tasks.Tasks {
	return tasks.New(dc)
}
