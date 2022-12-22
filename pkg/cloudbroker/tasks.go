package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/tasks"
)

// Accessing the tasks method group
func (cb *CloudBroker) Tasks() *tasks.Tasks {
	return tasks.New(cb.client)
}
